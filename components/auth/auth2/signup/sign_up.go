package signup

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"strings"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// AuthorizerForSignUp ...
type AuthorizerForSignUp struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	UserDAO  rbacdb.UserDAO         //starter:inject("#")
	PhoneDAO rbacdb.PhoneNumberDAO  //starter:inject("#")
	EmailDAO rbacdb.EmailAddressDAO //starter:inject("#")
	Agent    rbacdb.LocalAgent      //starter:inject("#")
}

func (inst *AuthorizerForSignUp) _impl() (auth.Registry, auth.Authorizer, auth.Mechanism) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *AuthorizerForSignUp) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:      0,
		Enabled:       true,
		Authenticator: nil,
		Authorizer:    inst,
		Mechanism:     inst,
	}
	return []*auth.Registration{r1}
}

// Support 。。。
func (inst *AuthorizerForSignUp) Support(r auth.Request) bool {
	action := r.Action()
	return action == auth.ActionSignUp
}

func (inst *AuthorizerForSignUp) getPassword(a auth.Authorization) (lang.Base64, error) {
	str, err := a.Parameters().GetParamRequired("password")
	if err != nil {
		return "", err
	}
	b64 := lang.Base64(str)
	return b64, nil
}

// Authorize ...
func (inst *AuthorizerForSignUp) Authorize(a auth.Authorization) error {

	ids := a.Identities()
	acc := &account{}

	password, err := inst.getPassword(a)
	if err != nil {
		return err
	}
	acc.password = password.Bytes()

	for _, id := range ids {
		inst.prepareAccountWithIdentity(acc, id)
	}
	inst.prepareAccountUser(acc)

	return inst.invokeTransaction(func(db *gorm.DB) error {
		return inst.apply(db, acc)
	})
}

func (inst *AuthorizerForSignUp) invokeTransaction(fn func(db *gorm.DB) error) error {
	db := inst.Agent.DB(nil)
	return db.Transaction(fn)
}

func (inst *AuthorizerForSignUp) getNamer(acc *account) *namer {
	n := &namer{}
	return n.init(acc)
}

func (inst *AuthorizerForSignUp) prepareAccountUser(acc *account) error {

	salt := make([]byte, 32)
	rand.Read(salt)

	pc := auth.PasswordCalculator{}
	pc.Reset().WithSalt(salt)
	sum := pc.Compute(acc.password)

	aNamer := inst.getNamer(acc)

	user := &rbacdb.UserEntity{}
	user.Password = lang.HexFromBytes(sum)
	user.Salt = lang.HexFromBytes(salt)
	user.Enabled = true
	user.Roles = rbac.NewRoleNameList(rbac.RoleUser)
	user.Nickname = aNamer.NickName()
	user.Name = aNamer.UserName()
	user.Avatar = "" // todo ...

	acc.user = user
	return nil
}

func (inst *AuthorizerForSignUp) prepareAccountWithIdentity(acc *account, id auth.Identity) error {
	mech := id.Mechanism()
	switch mech {
	case auth.MechanismSMS:
		xid, ok := id.(auth.PhoneIdentity)
		if ok {
			return inst.prepareAccountWithPhone(acc, xid)
		}
		break
	case auth.MechanismEmail:
		xid, ok := id.(auth.EmailIdentity)
		if ok {
			return inst.prepareAccountWithEmail(acc, xid)
		}
		break
	}
	return nil
}

func (inst *AuthorizerForSignUp) prepareAccountWithEmail(acc *account, id auth.EmailIdentity) error {

	addr := id.Address()

	mail := &rbacdb.EmailAddressEntity{}
	mail.Address = addr
	mail.Verified = true

	acc.mail = mail
	return nil
}

func (inst *AuthorizerForSignUp) prepareAccountWithPhone(acc *account, id auth.PhoneIdentity) error {

	full := id.FullNumber()
	pureNum := full.Pure()
	region, simpleNum, err := full.Parts()
	if err != nil {
		return err
	}

	phone := &rbacdb.PhoneNumberEntity{}
	phone.FullNumber = pureNum
	phone.SimpleNumber = simpleNum
	phone.Region = region
	phone.Verified = true

	acc.phone = phone
	return nil
}

func (inst *AuthorizerForSignUp) apply(db *gorm.DB, acc *account) error {

	user := acc.user
	mail := acc.mail
	phone := acc.phone

	user2, err := inst.UserDAO.Insert(db, user)
	if err != nil {
		return err
	}
	user = user2

	if phone != nil {
		phone.Owner = user.ID
		phone2, err := inst.PhoneDAO.Insert(db, phone)
		if err != nil {
			return err
		}
		phone = phone2
		user.Phone = phone2.ID
	}

	if mail != nil {
		mail.Owner = user.ID
		mail2, err := inst.EmailDAO.Insert(db, mail)
		if err != nil {
			return err
		}
		mail = mail2
		user.Email = mail2.ID
	}

	_, err = inst.UserDAO.Update(db, user.ID, func(obj *rbacdb.UserEntity) {
		obj.Email = user.Email
		obj.Phone = user.Phone
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////

type namer struct {
	seed string
}

func (inst *namer) init(acc *account) *namer {

	b := strings.Builder{}
	now := lang.Now()
	const nl = "\n"

	if acc.mail != nil {
		b.WriteString(acc.mail.Address.String())
	}

	if acc.phone != nil {
		b.WriteString(acc.phone.FullNumber.String())
	}

	b.WriteString(nl)
	b.WriteString(now.String())
	b.WriteString(nl)
	inst.seed = b.String()
	return inst
}

func (inst *namer) makeName(class string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(inst.seed)
	buffer.WriteString("\n" + class)
	data := buffer.Bytes()
	sum := sha1.Sum(data)
	hex := lang.HexFromBytes(sum[:])
	str := hex.String()
	return str[0:16]
}

func (inst *namer) UserName() rbac.UserName {
	str := "user-" + inst.makeName("user-name")
	return rbac.UserName(str)
}

func (inst *namer) NickName() string {
	str := "U" + inst.makeName("nick-name")
	return strings.ToUpper(str)
}

////////////////////////////////////////////////////////////////////////////////

type account struct {
	user     *rbacdb.UserEntity
	mail     *rbacdb.EmailAddressEntity
	phone    *rbacdb.PhoneNumberEntity
	password []byte
}
