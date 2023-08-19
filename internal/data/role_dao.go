package data

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// RoleDao ...
type RoleDao struct {
	//starter:component
	_as func(rbac.RoleDAO) //starter:as("#")

	Agent         libgorm.Agent      //starter:inject("#")
	UUIDGenerator lang.UUIDGenerator //starter:inject("#")
}

func (inst *RoleDao) _impl() rbac.RoleDAO {
	return inst
}

func (inst *RoleDao) model() *rbac.RoleEntity {
	return &rbac.RoleEntity{}
}

func (inst *RoleDao) modelList() []*rbac.RoleEntity {
	return make([]*rbac.RoleEntity, 0)
}

func (inst *RoleDao) makeResult(res *gorm.DB, o *rbac.RoleEntity) (*rbac.RoleEntity, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *RoleDao) Find(db *gorm.DB, id rbac.RoleID) (*rbac.RoleEntity, error) {
	db = inst.Agent.DB(db)
	o := inst.model()
	res := db.Find(o, id)
	return inst.makeResult(res, o)
}

// List ...
func (inst *RoleDao) List(db *gorm.DB, q *rbac.RoleQuery) ([]*rbac.RoleEntity, error) {

	if q == nil {
		q = &rbac.RoleQuery{}
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

// Insert ...
func (inst *RoleDao) Insert(db *gorm.DB, o *rbac.RoleEntity) (*rbac.RoleEntity, error) {

	uuid := inst.UUIDGenerator.Generate("rbac.RoleEntity")

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(res, o)
}

// Update ...
func (inst *RoleDao) Update(db *gorm.DB, id rbac.RoleID, updater func(*rbac.RoleEntity)) (*rbac.RoleEntity, error) {
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
func (inst *RoleDao) Delete(db *gorm.DB, id rbac.RoleID) error {
	db = inst.Agent.DB(db)
	res := db.Delete(inst.model(), id)
	return res.Error
}
