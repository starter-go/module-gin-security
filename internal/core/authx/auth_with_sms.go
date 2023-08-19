package authx

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/module-security-gin-gorm/internal/services"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/random"
	"github.com/starter-go/security/rbac"
)

// AuthWithSMS ...
type AuthWithSMS struct {

	//starter:component
	_as func(auth.Authenticator) //starter:as(".")

	RandomSer      random.Service                //starter:inject("#")
	MailSer        mails.Service                 //starter:inject("#")
	JWTSer         jwt.Service                   //starter:inject("#")
	VeriSer        services.VerificationService  //starter:inject("#")
	AuthorSer      services.AuthorizationService //starter:inject("#")
	DBAgent        libgorm.Agent                 //starter:inject("#")
	Users          rbac.UserDAO                  //starter:inject("#")
	EmailAddresses rbac.EmailAddressDAO          //starter:inject("#")
	PhoneNumbers   rbac.PhoneNumberDAO           //starter:inject("#")

}

func (inst *AuthWithSMS) _impl() (auth.Authenticator, auth.Authorizer, auth.Mechanism, auth.Registry) {
	return inst, inst, inst, inst
}

// ListRegistrations ...
func (inst *AuthWithSMS) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Enabled:       true,
		Priority:      0,
		Mechanism:     inst,
		Authenticator: inst,
		Authorizer:    inst,
	}
	return []*auth.Registration{r1}
}

// Support ...
func (inst *AuthWithSMS) Support(a auth.Request) bool {
	return a.Mechanism() == "sms"
}

// Authenticate ...
func (inst *AuthWithSMS) Authenticate(a auth.Authentication) error {
	x := inst.makeNewVericodeAuthx()
	return x.Authenticate(a)
}

// Authorize ...
func (inst *AuthWithSMS) Authorize(a auth.Authorization) error {
	x := inst.makeNewVericodeAuthx()
	return x.Authorize(a)
}

func (inst *AuthWithSMS) makeAddressForPhoneNum(num string) (mails.Address, error) {
	// src := []rune(num)
	// b := strings.Builder{}
	// for _, ch := range src {
	// 	if '0' <= ch && ch <= '9' {
	// 		b.WriteRune(ch)
	// 	}
	// }
	// if b.Len() < 10 {
	// 	return "", fmt.Errorf("bad phone number: %s", num)
	// }
	// b.WriteString("@phone")
	// str := b.String()
	// return mails.Address(str), nil
	return "", nil
}

func (inst *AuthWithSMS) makeNewVericodeAuthx() *vericodeAuthx {

	ctx := &vericodeContext{
		from:   "12345678@phone",
		length: 6,
	}

	ctx.RandomService = inst.RandomSer
	ctx.MailsService = inst.MailSer
	ctx.VerificationService = inst.VeriSer
	ctx.AuthorizationService = inst.AuthorSer
	ctx.dbAgent = inst.DBAgent
	ctx.users = inst.Users
	ctx.emailAddresses = inst.EmailAddresses
	ctx.phoneNumbers = inst.PhoneNumbers

	return &vericodeAuthx{context: ctx}
}
