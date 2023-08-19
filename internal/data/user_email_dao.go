package data

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// EmailDao ...
type EmailDao struct {

	//starter:component
	_as func(rbac.EmailAddressDAO) //starter:as("#")

	Agent         libgorm.Agent      //starter:inject("#")
	UUIDGenerator lang.UUIDGenerator //starter:inject("#")

}

func (inst *EmailDao) _impl() rbac.EmailAddressDAO {
	return inst
}

func (inst *EmailDao) model() *rbac.EmailAddressEntity {
	return &rbac.EmailAddressEntity{}
}

func (inst *EmailDao) modelList() []*rbac.EmailAddressEntity {
	return make([]*rbac.EmailAddressEntity, 0)
}

func (inst *EmailDao) makeResult(res *gorm.DB, o *rbac.EmailAddressEntity) (*rbac.EmailAddressEntity, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Insert ...
func (inst *EmailDao) Insert(db *gorm.DB, o *rbac.EmailAddressEntity) (*rbac.EmailAddressEntity, error) {

	uuid := inst.UUIDGenerator.Generate("rbac.EmailAddressEntity")

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(res, o)
}

// Update ...
func (inst *EmailDao) Update(db *gorm.DB, id rbac.EmailAddressID, updater func(*rbac.EmailAddressEntity)) (*rbac.EmailAddressEntity, error) {
	o := inst.model()
	db = inst.Agent.DB(db)
	res := db.Find(o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	updater(o)
	res = db.Save(o)
	return inst.makeResult(res, o)
}

// Delete ...
func (inst *EmailDao) Delete(db *gorm.DB, id rbac.EmailAddressID) error {
	db = inst.Agent.DB(db)
	res := db.Delete(inst.model(), id)
	return res.Error
}

// Find ...
func (inst *EmailDao) Find(db *gorm.DB, id rbac.EmailAddressID) (*rbac.EmailAddressEntity, error) {
	db = inst.Agent.DB(db)
	o := inst.model()
	res := db.Find(o, id)
	return inst.makeResult(res, o)
}

// List ...
func (inst *EmailDao) List(db *gorm.DB, q *rbac.EmailAddressQuery) ([]*rbac.EmailAddressEntity, error) {

	if q == nil {
		q = &rbac.EmailAddressQuery{}
		q.Pagination.Size = 10
	}

	offset := q.Pagination.Offset()
	limit := q.Pagination.Limit()
	list := inst.modelList()
	count := int64(0)

	db = inst.Agent.DB(db)
	db.Model(inst.model()).Count(&count)
	db = db.Offset(int(offset)).Limit(int(limit))
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	q.Pagination.Total = count
	return list, nil
}
