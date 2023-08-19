package authx

import (
	"fmt"
	"strconv"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/module-security-gin-gorm/internal/services"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
	"github.com/starter-go/vlog"
	"gorm.io/gorm"
)

// 这里是 auth_with_email 和 auth_with_sms 共享的代码

////////////////////////////////////////////////////////////////////////////////

type vericode struct {
	id      string
	account string
	code    string
	salt    lang.Hex
	sum     lang.Hex
}

////////////////////////////////////////////////////////////////////////////////

type vericodeContext struct {
	VerificationService  services.VerificationService
	AuthorizationService services.AuthorizationService
	RandomService        random.Service
	MailsService         mails.Service
	from                 mails.Address
	dbAgent              libgorm.Agent
	users                rbac.UserDAO
	emailAddresses       rbac.EmailAddressDAO
	phoneNumbers         rbac.PhoneNumberDAO

	length int // length of code
}

func (inst *vericodeContext) makeSalt() []byte {
	return inst.RandomService.NextBytes(512 / 8)
}

func (inst *vericodeContext) getNewPassword(r auth.Request) ([]byte, error) {
	const key = "password"
	raw, err := r.Parameters().GetParamRequired(key)
	if err != nil {
		return nil, err
	}
	b := lang.Base64(raw)
	data := b.Bytes()
	if len(data) < 1 {
		return nil, fmt.Errorf("bad password")
	}
	return data, nil
}

func (inst *vericodeContext) getUser(r auth.Request) (*rbac.UserEntity, error) {
	user1, err := r.Attributes().GetAttributeRequired("user")
	if err != nil {
		return nil, err
	}
	user2, ok := user1.(*rbac.UserEntity)
	if !ok {
		return nil, fmt.Errorf("cannot cast to *rbac.UserEntity")
	}
	return user2, nil
}

////////////////////////////////////////////////////////////////////////////////

type vericodeGenerator struct {
	context *vericodeContext
}

func (inst *vericodeGenerator) Generate(vc *vericode) error {

	id := inst.makeID()
	code := inst.makeCode()
	salt := inst.context.makeSalt()
	sum := make([]byte, 0)

	pc := auth.PasswordCalculator{}
	pc.Init(sum, salt)
	sum = pc.Compute([]byte(code))

	vc.id = id
	vc.code = code
	vc.salt = lang.HexFromBytes(salt)
	vc.sum = lang.HexFromBytes(sum)
	return nil
}

func (inst *vericodeGenerator) makeCode() string {
	code := ""
	length := inst.context.length
	if length == 0 {
		length = 6 // default value
	}
	for timeout := 32; timeout > 0; timeout-- {
		n := inst.context.RandomService.NextInt64()
		if n < 0 {
			n = -n
		}
		code = strconv.FormatInt(n, 10)
		if len(code) > length {
			code = code[0:length]
			break
		}
	}
	return code
}

func (inst *vericodeGenerator) makeID() string {
	b := inst.context.RandomService.NextBytes(200 / 8)
	uuid := lang.NewUUID(b)
	return uuid.String()
}

////////////////////////////////////////////////////////////////////////////////

type vericodeVerifier struct {
	context *vericodeContext
}

func (inst *vericodeVerifier) Verify(vc *vericode) error {
	salt := vc.salt.Bytes()
	sum := vc.sum.Bytes()
	plain := vc.code
	pc := auth.PasswordCalculator{}
	pc.Init(sum, salt)
	return pc.Verify([]byte(plain))
}

////////////////////////////////////////////////////////////////////////////////

type vericodeAuthx struct {
	context *vericodeContext
}

func (inst *vericodeAuthx) verify(a auth.Authentication) error {

	account := a.Account()
	code := a.Secret()
	uuid1 := a.Parameters().GetParamOptional("uuid", "")

	rec := inst.context.VerificationService.GetRecord(account)
	err := rec.Prepare(a.Action())
	if err != nil {
		return err
	}
	user, err := rec.Read()
	if err != nil {
		return err
	}

	uuid2 := user.UUID.String()
	if uuid1 != uuid2 {
		return fmt.Errorf("bad verification")
	}

	vc := &vericode{}
	vc.id = uuid1
	vc.account = account
	vc.code = string(code)
	vc.salt = user.Salt
	vc.sum = user.Password

	verifier := &vericodeVerifier{context: inst.context}
	return verifier.Verify(vc)
}

func (inst *vericodeAuthx) login(a auth.Authorization) error {
	ctx := a.Context()
	user, err := inst.context.getUser(a)
	if err != nil {
		return err
	}
	ser := inst.context.AuthorizationService
	return ser.Authorize(ctx, user)
}

func (inst *vericodeAuthx) signUp(a auth.Authorization) error {
	mechanism := a.Mechanism()
	switch mechanism {
	case auth.MechanismEmail:
		h := &handlerForSignUpWithEmail{context: inst.context}
		return h.execute(a)
	case auth.MechanismPhone:
		h := &handlerForSignUpWithPhone{context: inst.context}
		return h.execute(a)
	}
	return fmt.Errorf("bad mechanism for sign-up, mechanism=" + mechanism)
}

func (inst *vericodeAuthx) resetPassword(a auth.Authorization) error {
	h := &handlerForResetPassword{context: inst.context}
	return h.execute(a)
}

func (inst *vericodeAuthx) sendCode(a auth.Authentication) error {

	account := a.Account()

	rec := inst.context.VerificationService.GetRecord(account)
	err := rec.Prepare(a.Action())
	if err != nil {
		return err
	}

	// gen code
	generator := &vericodeGenerator{context: inst.context}
	vc := &vericode{}
	vc.account = account
	err = generator.Generate(vc)
	if err != nil {
		return err
	}

	// send message
	err = inst.sendMessageWithCode(a, vc)
	if err != nil {
		return err
	}

	// write to resp
	a.Parameters().SetParam("uuid", vc.id)

	// write to verification service
	user := &rbac.UserEntity{}
	user.Password = vc.sum
	user.Salt = vc.salt
	user.Name = rbac.UserName(account)
	user.UUID = lang.ParseUUID(vc.id)
	return rec.Write(user)
}

func (inst *vericodeAuthx) sendMessageWithCode(a auth.Authentication, vc *vericode) error {

	code := vc.code
	text := "Your verification code: " + code
	to := a.Account() // email-address | phone-number
	addr2 := mails.Address(to)

	msg := &mails.Message{}
	msg.Title = "Verification Code"
	msg.FromAddress = inst.context.from
	msg.ToAddresses = []mails.Address{addr2}
	msg.ContentType = "text/plain"
	msg.Content = []byte(text)

	logger := vlog.GetLogger()
	if logger.IsDebugEnabled() {
		logger.Debug("msg.content = %s", text)
	}

	ctx := a.Context()
	return inst.context.MailsService.Send(ctx, msg)
}

// Authenticate 验证
func (inst *vericodeAuthx) Authenticate(a auth.Authentication) error {
	action := a.Action()
	switch action {

	case auth.ActionLogin:
		return inst.verify(a)

	case auth.ActionResetPassword:
		return inst.verify(a)

	case auth.ActionSignUp:
		return inst.verify(a)

	case auth.ActionSendCode:
		return inst.sendCode(a)
	}
	return fmt.Errorf("bad verification code")
}

// Authorize 授权
func (inst *vericodeAuthx) Authorize(a auth.Authorization) error {
	action := a.Action()
	switch action {

	case auth.ActionLogin:
		return inst.login(a)

	case auth.ActionResetPassword:
		return inst.resetPassword(a)

	case auth.ActionSignUp:
		return inst.signUp(a)

	case auth.ActionSendCode:
		return nil
	}
	return fmt.Errorf("bad verification auth action: %s", action)
}

////////////////////////////////////////////////////////////////////////////////

type handlerForSignUpWithEmail struct {
	context *vericodeContext
}

func (inst *handlerForSignUpWithEmail) execute(a auth.Authorization) error {

	db := inst.context.dbAgent.DB(nil)
	err := db.Transaction(func(tx *gorm.DB) error {
		return inst.execute2(tx, a)
	})
	return err // fmt.Errorf("no impl")
}

func (inst *handlerForSignUpWithEmail) execute2(db *gorm.DB, a1 auth.Authorization) error {

	a2 := a1.(auth.Authentication)

	addr := a2.Account() // email address
	password, err := inst.context.getNewPassword(a1)
	if err != nil {
		return err
	}

	user := &rbac.UserEntity{}
	user.Email = addr
	user.Roles = rbac.RoleNameList(rbac.RoleUser)

	inst.preparePassword(user, password)
	inst.prepareUserName(user)

	user2, err := inst.context.users.Insert(db, user)
	if err != nil {
		return err
	}
	user = user2

	email := &rbac.EmailAddressEntity{}
	email.Owner = user.ID
	email.Address = addr

	_, err = inst.context.emailAddresses.Insert(db, email)
	return err
}

func (inst *handlerForSignUpWithEmail) prepareUserName(user *rbac.UserEntity) error {

	now := lang.Now()
	n := now.Int()
	name := fmt.Sprintf("user-%d", n)

	user.Name = rbac.UserName(name)
	user.Nickname = name

	return nil
}

func (inst *handlerForSignUpWithEmail) preparePassword(user *rbac.UserEntity, password []byte) error {

	salt := inst.context.makeSalt()
	plain := password

	pc := &auth.PasswordCalculator{}
	pc.Init(nil, salt)
	sum := pc.Compute(plain)

	user.Password = lang.HexFromBytes(sum)
	user.Salt = lang.HexFromBytes(salt)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type handlerForSignUpWithPhone struct {
	context *vericodeContext
}

func (inst *handlerForSignUpWithPhone) execute(a auth.Authorization) error {
	return fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type handlerForResetPassword struct {
	context *vericodeContext
}

func (inst *handlerForResetPassword) execute(a auth.Authorization) error {
	salt := inst.context.makeSalt()
	password, err := inst.context.getNewPassword(a)
	if err != nil {
		return err
	}

	user, err := inst.context.getUser(a)
	if err != nil {
		return err
	}

	uid := user.ID
	pc := &auth.PasswordCalculator{}
	pc.Init(nil, salt)
	sum := pc.Compute(password)
	_, err = inst.context.users.Update(nil, uid, func(o *rbac.UserEntity) {
		o.Password = lang.HexFromBytes(sum)
		o.Salt = lang.HexFromBytes(salt)
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////
