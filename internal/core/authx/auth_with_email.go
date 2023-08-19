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

// AuthWithEmail 。。。
type AuthWithEmail struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

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

func (inst *AuthWithEmail) _impl() auth.Registry {
	return inst
}

// ListRegistrations ...
func (inst *AuthWithEmail) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:      0,
		Enabled:       true,
		Authenticator: inst,
		Authorizer:    inst,
		Mechanism:     inst,
	}
	return []*auth.Registration{r1}
}

// Authorize 授权
func (inst *AuthWithEmail) Authorize(a auth.Authorization) error {
	x := inst.makeNewVericodeAuthx()
	return x.Authorize(a)
}

// Authenticate 验证
func (inst *AuthWithEmail) Authenticate(a auth.Authentication) error {
	x := inst.makeNewVericodeAuthx()
	return x.Authenticate(a)
}

// Support ...
func (inst *AuthWithEmail) Support(a auth.Request) bool {
	return a.Mechanism() == auth.MechanismEmail
}

func (inst *AuthWithEmail) makeNewVericodeAuthx() *vericodeAuthx {

	ctx := &vericodeContext{
		from:   "mock@example.com",
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

////////////////////////////////////////////////////////////////////////////////
