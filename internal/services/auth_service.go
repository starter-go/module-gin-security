package services

import (
	"context"

	"github.com/starter-go/security/rbac"
)

// // AuthService 验证服务
// type AuthService interface {
// 	rbac.AuthService

// 	// 注册新用户
// 	AddUser(ctx context.Context, o *rbac.AuthDTO) (*rbac.AuthDTO, error)

// 	// 改密码
// 	ChangePassword(ctx context.Context, o *rbac.AuthDTO) (*rbac.AuthDTO, error)

// 	// 重设密码
// 	ResetPassword(ctx context.Context, o *rbac.AuthDTO) (*rbac.AuthDTO, error)

// 	// 发送验证码（通过 email 或 sms）
// 	SendCode(ctx context.Context, o *rbac.AuthDTO) (*rbac.AuthDTO, error)
// }

// AuthorizationService 授权服务
type AuthorizationService interface {

	// 向用户授权
	Authorize(ctx context.Context, u *rbac.UserEntity) error
}
