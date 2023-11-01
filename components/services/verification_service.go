package services

import (
	"context"

	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
)

// VerificationRecord ...
type VerificationRecord interface {
	Account() string

	// 使用 UserEntity 来保存验证记录
	Read() (*rbacdb.UserEntity, error)

	// 使用 UserEntity 来保存验证记录
	Write(u *rbacdb.UserEntity) error

	Prepare(action string) error
}

// Verification ... 验证码的相关信息
type Verification struct {
	Mechanism string
	Code      string
	ToMail    rbac.EmailAddress    // 根据 Mechanism 取值
	ToPhone   rbac.FullPhoneNumber // 根据 Mechanism 取值
}

// VerificationService ...
type VerificationService interface {
	// GetRecord(account string) VerificationRecord

	Verify(c context.Context, v *Verification) error

	SendCode(c context.Context, v *Verification) error
}
