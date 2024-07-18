package login

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/subjects"
)

// AuthorizerForLogin ...
type AuthorizerForLogin struct {

	//starter:component
	_as func(auth.Registry) //starter:as(".")

	// UserDAO  rbacdb.UserDAO         //starter:inject("#")
	// PhoneDAO rbacdb.PhoneNumberDAO  //starter:inject("#")
	// EmailDAO rbacdb.EmailAddressDAO //starter:inject("#")
	// Agent    rbacdb.LocalAgent      //starter:inject("#")

	JWT jwt.Service //starter:inject("#")

}

func (inst *AuthorizerForLogin) _impl() (auth.Registry, auth.Authorizer, auth.Mechanism) {
	return inst, inst, inst
}

// ListRegistrations ...
func (inst *AuthorizerForLogin) ListRegistrations() []*auth.Registration {
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
func (inst *AuthorizerForLogin) Support(r auth.Request) bool {
	action := r.Action()
	return action == auth.ActionLogin
}

// Authorize ...
func (inst *AuthorizerForLogin) Authorize(a auth.Authorization) error {
	ids := a.Identities()
	for _, id := range ids {
		ok, err := inst.handle(id)
		if err != nil {
			return err
		} else if ok {
			return nil
		}
	}
	return fmt.Errorf("no user info to login")
}

func (inst *AuthorizerForLogin) handle(identity auth.Identity) (bool, error) {

	uid, ok := identity.(auth.UserIdentity)
	if !ok {
		return false, nil
	}

	ctx := identity.By().Context()
	user := &rbac.UserDTO{}
	user.ID = uid.UserID()
	user.Avatar = uid.Avatar()
	user.Name = uid.UserName()
	user.NickName = uid.Nickname()
	user.Roles = uid.Roles()

	err := inst.createNewSession(ctx, user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *AuthorizerForLogin) createNewSession(c context.Context, user *rbac.UserDTO) error {

	if user == nil {
		return fmt.Errorf("no user info")
	}

	sub, err := subjects.Current(c)
	if err != nil {
		return err
	}

	now := lang.Now()
	tokenRef := sub.GetToken()
	sessionRef := sub.GetSession()

	// init current.User
	cu := &rbac.CurrentUser{
		User:     user.ID,
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		Roles:    user.Roles,
	}

	// init session
	session := sessionRef.Get()
	session.CurrentUser = *cu
	session.Authenticated = true
	session.CreatedAt = now
	session.UpdatedAt = now
	session.StartedAt = now
	session.ExpiredAt = inst.sessionExpiredAt(now)

	// create session
	sessionRef.Create()
	sessid := sessionRef.SessionID()
	session.ID = sessid
	session.UUID = lang.UUID(sessid)

	// init token
	token := tokenRef.Get()
	token.CurrentUser = *cu
	token.Session = sessid
	token.StartedAt = now
	token.ExpiredAt = inst.tokenExpiredAt(now)

	// commit
	tokenRef.Set(token)
	sessionRef.Set(session)
	err = sessionRef.Commit()
	if err != nil {
		return err
	}
	return tokenRef.Commit()
}

func (inst *AuthorizerForLogin) sessionExpiredAt(t0 lang.Time) lang.Time {
	// 180 min
	const delta = 180 * 60 * 1000
	return t0 + delta
}

func (inst *AuthorizerForLogin) tokenExpiredAt(t0 lang.Time) lang.Time {
	// 10-min
	const delta = 10 * 60 * 1000
	return t0 + delta
}
