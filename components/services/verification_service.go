package services

import "github.com/starter-go/security-gorm/rbacdb"

// VerificationRecord ...
type VerificationRecord interface {
	Account() string

	// 使用 UserEntity 来保存验证记录
	Read() (*rbacdb.UserEntity, error)

	// 使用 UserEntity 来保存验证记录
	Write(u *rbacdb.UserEntity) error

	Prepare(action string) error
}

// VerificationService ...
type VerificationService interface {
	GetRecord(account string) VerificationRecord
}
