package services

import "github.com/starter-go/security/rbac"

// VerificationRecord ...
type VerificationRecord interface {
	Account() string

	// 使用 UserEntity 来保存验证记录
	Read() (*rbac.UserEntity, error)

	// 使用 UserEntity 来保存验证记录
	Write(u *rbac.UserEntity) error

	Prepare(action string) error
}

// VerificationService ...
type VerificationService interface {
	GetRecord(account string) VerificationRecord
}
