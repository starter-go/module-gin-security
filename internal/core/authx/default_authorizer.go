package authx

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/module-security-gin-gorm/internal/services"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/rbac"
)

// DefaultAuthorizer 是默认的授权组件
type DefaultAuthorizer struct {

	//starter:component
	_as func(services.AuthorizationService) //starter:as("#")

	Service       jwt.Service //starter:inject("#")
	TokenMaxAge   int         //starter:inject("${security.jwt.max-age}")
	SessionMaxAge int         //starter:inject("${security.session.max-age}")
}

func (inst *DefaultAuthorizer) _impl() services.AuthorizationService {
	return inst
}

// Authorize ...
func (inst *DefaultAuthorizer) Authorize(ctx context.Context, user *rbac.UserEntity) error {

	now := lang.Now()

	o1 := &jwt.Token{}
	user2 := &o1.Session.User

	user2.Avatar = user.Avatar
	user2.ID = user.ID
	user2.Name = user.Name
	user2.NickName = user.Nickname
	user2.Roles = user.Roles

	o1.CreatedAt = now
	o1.ExpiredAt = inst.computeExpiredAt(now, inst.TokenMaxAge)

	o1.Session.CreatedAt = now
	o1.Session.ExpiredAt = inst.computeExpiredAt(now, inst.SessionMaxAge)
	o1.Session.Authorized = true

	t1, err := inst.Service.Encode(o1)
	if err != nil {
		return err
	}
	return inst.Service.SetText(ctx, t1)
}

func (inst *DefaultAuthorizer) computeExpiredAt(now lang.Time, maxAge int) lang.Time {
	ma := lang.Time(maxAge)
	if ma < 0 {
		ma = 0
	}
	return now + ma
}
