package data

import (
	"fmt"

	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// PhoneDao ...
type PhoneDao struct {

	//starter:component
	_as func(rbac.PhoneNumberDAO) //starter:as("#")

}

func (inst *PhoneDao) _impl() rbac.PhoneNumberDAO {
	return inst
}

// Insert ...
func (inst *PhoneDao) Insert(db *gorm.DB, o *rbac.PhoneNumberEntity) (*rbac.PhoneNumberEntity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *PhoneDao) Update(db *gorm.DB, id rbac.PhoneNumberID, updater func(*rbac.PhoneNumberEntity)) (*rbac.PhoneNumberEntity, error) {
	return nil, fmt.Errorf("no impl")
}

// Delete ...
func (inst *PhoneDao) Delete(db *gorm.DB, id rbac.PhoneNumberID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *PhoneDao) Find(db *gorm.DB, id rbac.PhoneNumberID) (*rbac.PhoneNumberEntity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *PhoneDao) List(db *gorm.DB, q *rbac.PhoneNumberQuery) ([]*rbac.PhoneNumberEntity, error) {
	return nil, fmt.Errorf("no impl")
}
