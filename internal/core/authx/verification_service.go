package authx

import (
	"fmt"
	"sync"

	"github.com/starter-go/module-security-gin-gorm/internal/services"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// VerificationServiceImpl ...
type VerificationServiceImpl struct {

	//starter:component
	_as func(services.VerificationService) //starter:as("#")

	mutex sync.Mutex
	table map[string]services.VerificationRecord
}

func (inst *VerificationServiceImpl) _impl() services.VerificationService {
	return inst
}

func (inst *VerificationServiceImpl) getTable() map[string]services.VerificationRecord {
	t := inst.table
	if t == nil {
		t = make(map[string]services.VerificationRecord)
		inst.table = t
	}
	return t
}

// GetRecord ...
func (inst *VerificationServiceImpl) GetRecord(account string) services.VerificationRecord {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	tab := inst.getTable()
	rec := tab[account]

	if rec == nil {
		vr := &verificationRecord{
			service: inst,
			account: account,
		}
		vr.entity.Name = rbac.UserName(account)
		rec = vr
	}

	return rec
}

func (inst *VerificationServiceImpl) invoke(fn func()) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()
	fn()
}

////////////////////////////////////////////////////////////////////////////////

type verificationRecord struct {
	service *VerificationServiceImpl
	account string
	entity  rbac.UserEntity
}

func (inst *verificationRecord) Account() string {
	return inst.account
}

// 使用 UserEntity 来保存验证记录
func (inst *verificationRecord) Read() (*rbac.UserEntity, error) {
	dst := &rbac.UserEntity{}
	inst.service.invoke(func() {
		*dst = inst.entity
	})
	return dst, nil
}

// 使用 UserEntity 来保存验证记录
func (inst *verificationRecord) Write(src *rbac.UserEntity) error {
	if src == nil {
		return fmt.Errorf("param is nil")
	}
	inst.service.invoke(func() {
		tab := inst.service.getTable()
		key := inst.account
		inst.entity = *src
		tab[key] = inst
	})
	return nil
}

func (inst *verificationRecord) Prepare(action string) error {

	switch action {
	case auth.ActionSendCode:
	default:
	}

	// todo ...
	return nil
}
