package data

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// UserDao ...
type UserDao struct {
	//starter:component
	_as func(rbac.UserDAO) //starter:as("#")

	Agent         libgorm.Agent      //starter:inject("#")
	UUIDGenerator lang.UUIDGenerator //starter:inject("#")
}

func (inst *UserDao) _impl() {
	inst._as(inst)
}

func (inst *UserDao) model() *rbac.UserEntity {
	return &rbac.UserEntity{}
}

func (inst *UserDao) modelList() []*rbac.UserEntity {
	return make([]*rbac.UserEntity, 0)
}

func (inst *UserDao) makeResult(res *gorm.DB, o *rbac.UserEntity) (*rbac.UserEntity, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *UserDao) Find(db *gorm.DB, id rbac.UserID) (*rbac.UserEntity, error) {
	db = inst.Agent.DB(db)
	o := inst.model()
	res := db.Find(o, id)
	return inst.makeResult(res, o)
}

// List ...
func (inst *UserDao) List(db *gorm.DB, q *rbac.UserQuery) ([]*rbac.UserEntity, error) {

	if q == nil {
		q = &rbac.UserQuery{}
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
func (inst *UserDao) Insert(db *gorm.DB, o *rbac.UserEntity) (*rbac.UserEntity, error) {

	uuid := inst.UUIDGenerator.Generate("rbac.UserEntity")

	o.ID = 0
	o.UUID = uuid

	db = inst.Agent.DB(db)
	res := db.Create(o)
	return inst.makeResult(res, o)
}

// Update ...
func (inst *UserDao) Update(db *gorm.DB, id rbac.UserID, updater func(*rbac.UserEntity)) (*rbac.UserEntity, error) {
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
func (inst *UserDao) Delete(db *gorm.DB, id rbac.UserID) error {
	db = inst.Agent.DB(db)
	res := db.Delete(inst.model(), id)
	return res.Error
}

// FindByName ...
func (inst *UserDao) FindByName(db *gorm.DB, text string) (*rbac.UserEntity, error) {
	o := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("name = ?", text).First(o)
	return inst.makeResult(res, o)
}

// FindByEmail ...
func (inst *UserDao) FindByEmail(db *gorm.DB, text string) (*rbac.UserEntity, error) {
	o := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("email = ?", text).First(o)
	return inst.makeResult(res, o)
}

// FindByPhone ...
func (inst *UserDao) FindByPhone(db *gorm.DB, text string) (*rbac.UserEntity, error) {
	o := inst.model()
	db = inst.Agent.DB(db)
	res := db.Where("phone = ?", text).First(o)
	return inst.makeResult(res, o)
}

// FindByAny ...
func (inst *UserDao) FindByAny(db *gorm.DB, text string) (*rbac.UserEntity, error) {

	if text == "" {
		return nil, fmt.Errorf("no user")
	}

	ent, err := inst.FindByName(db, text)
	if err == nil {
		return ent, nil
	}

	ent, err = inst.FindByEmail(db, text)
	if err == nil {
		return ent, nil
	}

	ent, err = inst.FindByPhone(db, text)
	if err == nil {
		return ent, nil
	}

	return nil, fmt.Errorf("no user")
}
