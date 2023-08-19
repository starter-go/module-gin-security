package authx

import (
	"fmt"
	"strings"

	"github.com/starter-go/module-security-gin-gorm/internal/services"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// PasswordAuth ...
type PasswordAuth struct {

	//starter:component
	_as func(auth.Authenticator, auth.Registry) //starter:as(".",".")

	UserDao rbac.UserDAO                  //starter:inject("#")
	AuthSer services.AuthorizationService //starter:inject("#")
}

func (inst *PasswordAuth) _impl() (auth.Authenticator, auth.Authorizer, auth.Registry) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *PasswordAuth) ListRegistrations() []*auth.Registration {
	r1 := &auth.Registration{
		Priority:      0,
		Enabled:       true,
		Authenticator: inst,
		Authorizer:    inst,
		Mechanism:     inst,
	}
	return []*auth.Registration{r1}
}

// Support ...
func (inst *PasswordAuth) Support(a auth.Request) bool {
	mech := a.Mechanism()
	mech = strings.TrimSpace(mech)
	mech = strings.ToLower(mech)
	return (mech == "password")
}

// Authenticate ...
func (inst *PasswordAuth) Authenticate(a auth.Authentication) error {

	account := a.Account()
	user1, err := inst.UserDao.FindByAny(nil, account)
	if err != nil {
		return err
	}

	target := user1.Password.Bytes()
	salt := user1.Salt.Bytes()
	pc := auth.PasswordCalculator{}
	pc.Init(target, salt)
	err = pc.Verify(a.Secret())
	if err != nil {
		return err
	}

	a.Attributes().SetAttribute("user", user1)
	return nil
}

// Authorize ...
func (inst *PasswordAuth) Authorize(a auth.Authorization) error {

	action := a.Action()

	if action == "" || action == "login" {
		return inst.doLogin(a)
	}

	return fmt.Errorf("bad auth.action:" + action)
}

func (inst *PasswordAuth) doLogin(a auth.Authorization) error {

	ctx := a.Context()
	user1, err := a.Attributes().GetAttributeRequired("user")
	if err != nil {
		return err
	}

	user2, ok := user1.(*rbac.UserEntity)
	if !ok {
		return fmt.Errorf("bad cast to *rbac.UserEntity")
	}

	return inst.AuthSer.Authorize(ctx, user2)
}
