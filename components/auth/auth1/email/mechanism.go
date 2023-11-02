package email

import (
	"strings"

	"github.com/starter-go/security-gin-gorm/components/services"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// AuthByEmail  提供邮件验证功能
type AuthByEmail struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	VerificationService services.VerificationService //starter:inject("#")

}

func (inst *AuthByEmail) _impl() auth.Registry {
	return inst
}

// ListRegistrations ...
func (inst *AuthByEmail) ListRegistrations() []*auth.Registration {

	a1 := &authenticator{parent: inst}
	a2 := &authorizer{}

	r1 := &auth.Registration{
		Authenticator: a1,
		Authorizer:    nil,
		Mechanism:     a1,
		Priority:      0,
		Enabled:       true,
	}
	r2 := &auth.Registration{
		Authenticator: nil,
		Authorizer:    a2,
		Mechanism:     a2,
		Priority:      0,
		Enabled:       true,
	}
	return []*auth.Registration{r1, r2}
}

////////////////////////////////////////////////////////////////////////////////

type authenticator struct {
	parent *AuthByEmail
}

func (inst *authenticator) _impl() (auth.Authenticator, auth.Mechanism) {
	return inst, inst
}

func (inst *authenticator) Support(r auth.Request) bool {
	a, ok := r.(auth.Authentication)
	if !ok {
		return false
	}
	mech := a.Mechanism()
	return mech == auth.MechanismEmail
}

func (inst *authenticator) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	action := a.Action()
	if action == auth.ActionSendCode {
		return inst.sendCode(a)
	}
	return inst.verify(a)
}

func (inst *authenticator) sendCode(a auth.Authentication) ([]auth.Identity, error) {
	// 发送验证邮件
	addr := a.Account()
	ctx := a.Context()
	ser := inst.parent.VerificationService
	veri := &services.Verification{
		Mechanism: a.Mechanism(),
		ToMail:    rbac.EmailAddress(addr),
	}
	err := ser.SendCode(ctx, veri)
	if err != nil {
		return nil, err
	}
	ids := make([]auth.Identity, 0)
	return ids, nil
}

func (inst *authenticator) verify(a auth.Authentication) ([]auth.Identity, error) {
	// 验证 code
	addr := a.Account()
	ctx := a.Context()
	ser := inst.parent.VerificationService
	code := a.Secret()
	veri := &services.Verification{
		Mechanism: a.Mechanism(),
		ToMail:    rbac.EmailAddress(addr),
		Code:      string(code),
	}
	err := ser.Verify(ctx, veri)
	if err != nil {
		return nil, err
	}
	return inst.makeResultOK(a)
}

func (inst *authenticator) makeResultOK(a auth.Authentication) ([]auth.Identity, error) {

	addr := a.Account()
	addr = strings.TrimSpace(addr)

	info := &rbac.EmailAddressDTO{}
	info.Address = rbac.EmailAddress(addr)

	id := auth.NewEmailIdentity(a, info)
	ids := make([]auth.Identity, 0)
	ids = append(ids, id)
	return ids, nil
}

////////////////////////////////////////////////////////////////////////////////

type authorizer struct{}

func (inst *authorizer) _impl() (auth.Authorizer, auth.Mechanism) {
	return inst, inst
}

func (inst *authorizer) Support(r auth.Request) bool {
	a, ok := r.(auth.Authorization)
	if !ok {
		return false
	}
	return a.Action() == auth.ActionSendCode
}

func (inst *authorizer) Authorize(a auth.Authorization) error {
	// NOP : 这个授权组件仅仅用于在发送 email 时占位，不实现任何实际的功能
	return nil
}

////////////////////////////////////////////////////////////////////////////////
