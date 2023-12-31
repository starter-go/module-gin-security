package sms

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/services"
	"github.com/starter-go/security/auth"
)

// AuthBySMS  提供邮件验证功能
type AuthBySMS struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	VerificationService services.VerificationService //starter:inject("#")

}

func (inst *AuthBySMS) _impl() auth.Registry {
	return inst
}

// ListRegistrations ...
func (inst *AuthBySMS) ListRegistrations() []*auth.Registration {

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
	parent *AuthBySMS
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
	return mech == auth.MechanismSMS
}

func (inst *authenticator) Authenticate(a auth.Authentication) ([]auth.Identity, error) {
	action := a.Action()
	if action == auth.ActionSendCode {
		return inst.sendCode(a)
	}
	return inst.verify(a)
}

func (inst *authenticator) sendCode(a auth.Authentication) ([]auth.Identity, error) {
	// 发送验证短信
	num, err := rbac.ParseFullPhoneNumber(a.Account())
	if err != nil {
		return nil, err
	}
	ctx := a.Context()
	ser := inst.parent.VerificationService
	veri := &services.Verification{
		Mechanism: a.Mechanism(),
		ToPhone:   num,
	}
	err = ser.SendCode(ctx, veri)
	if err != nil {
		return nil, err
	}
	ids := make([]auth.Identity, 0)
	return ids, nil
}

func (inst *authenticator) verify(a auth.Authentication) ([]auth.Identity, error) {
	// 验证 code
	num, err := rbac.ParseFullPhoneNumber(a.Account())
	if err != nil {
		return nil, err
	}
	ctx := a.Context()
	ser := inst.parent.VerificationService
	code := a.Secret()
	veri := &services.Verification{
		Mechanism: a.Mechanism(),
		ToPhone:   num,
		Code:      string(code),
	}
	err = ser.Verify(ctx, veri)
	if err != nil {
		return nil, err
	}
	return inst.makeResultOK(a)
}

func (inst *authenticator) makeResultOK(a auth.Authentication) ([]auth.Identity, error) {

	full, err := rbac.ParseFullPhoneNumber(a.Account())
	if err != nil {
		return nil, err
	}
	region, phone, err := full.Parts()
	if err != nil {
		return nil, err
	}

	info := &rbac.PhoneNumberDTO{}
	info.FullNumber = full
	info.RegionCode2 = region
	info.SimpleNumber = phone

	id := auth.NewPhoneIdentity(a, info)
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
	// NOP : 这个授权组件仅仅用于在发送 sms 时占位，不实现任何实际的功能
	return nil
}

////////////////////////////////////////////////////////////////////////////////
