package gen4data
import (
    p3f0a0f8fe "github.com/starter-go/base/lang"
    p512a30914 "github.com/starter-go/libgorm"
    p4879aa5e6 "github.com/starter-go/module-security-gin-gorm/internal/data"
     "github.com/starter-go/application"
)

// type p4879aa5e6.PermissionDao in package:github.com/starter-go/module-security-gin-gorm/internal/data
//
// id:com-4879aa5e6a810f44-data-PermissionDao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO
// scope:singleton
//
type p4879aa5e6a_data_PermissionDao struct {
}

func (inst* p4879aa5e6a_data_PermissionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4879aa5e6a810f44-data-PermissionDao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4879aa5e6a_data_PermissionDao) new() any {
    return &p4879aa5e6.PermissionDao{}
}

func (inst* p4879aa5e6a_data_PermissionDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4879aa5e6.PermissionDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*p4879aa5e6a_data_PermissionDao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p4879aa5e6a_data_PermissionDao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type p4879aa5e6.RoleDao in package:github.com/starter-go/module-security-gin-gorm/internal/data
//
// id:com-4879aa5e6a810f44-data-RoleDao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO
// scope:singleton
//
type p4879aa5e6a_data_RoleDao struct {
}

func (inst* p4879aa5e6a_data_RoleDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4879aa5e6a810f44-data-RoleDao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4879aa5e6a_data_RoleDao) new() any {
    return &p4879aa5e6.RoleDao{}
}

func (inst* p4879aa5e6a_data_RoleDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4879aa5e6.RoleDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*p4879aa5e6a_data_RoleDao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p4879aa5e6a_data_RoleDao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type p4879aa5e6.UserDao in package:github.com/starter-go/module-security-gin-gorm/internal/data
//
// id:com-4879aa5e6a810f44-data-UserDao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserDAO
// scope:singleton
//
type p4879aa5e6a_data_UserDao struct {
}

func (inst* p4879aa5e6a_data_UserDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4879aa5e6a810f44-data-UserDao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4879aa5e6a_data_UserDao) new() any {
    return &p4879aa5e6.UserDao{}
}

func (inst* p4879aa5e6a_data_UserDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4879aa5e6.UserDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*p4879aa5e6a_data_UserDao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p4879aa5e6a_data_UserDao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type p4879aa5e6.EmailDao in package:github.com/starter-go/module-security-gin-gorm/internal/data
//
// id:com-4879aa5e6a810f44-data-EmailDao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-EmailAddressDAO
// scope:singleton
//
type p4879aa5e6a_data_EmailDao struct {
}

func (inst* p4879aa5e6a_data_EmailDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4879aa5e6a810f44-data-EmailDao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-EmailAddressDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4879aa5e6a_data_EmailDao) new() any {
    return &p4879aa5e6.EmailDao{}
}

func (inst* p4879aa5e6a_data_EmailDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4879aa5e6.EmailDao)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGenerator = inst.getUUIDGenerator(ie)


    return nil
}


func (inst*p4879aa5e6a_data_EmailDao) getAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p4879aa5e6a_data_EmailDao) getUUIDGenerator(ie application.InjectionExt)p3f0a0f8fe.UUIDGenerator{
    return ie.GetComponent("#alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator").(p3f0a0f8fe.UUIDGenerator)
}



// type p4879aa5e6.PhoneDao in package:github.com/starter-go/module-security-gin-gorm/internal/data
//
// id:com-4879aa5e6a810f44-data-PhoneDao
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PhoneNumberDAO
// scope:singleton
//
type p4879aa5e6a_data_PhoneDao struct {
}

func (inst* p4879aa5e6a_data_PhoneDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4879aa5e6a810f44-data-PhoneDao"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PhoneNumberDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4879aa5e6a_data_PhoneDao) new() any {
    return &p4879aa5e6.PhoneDao{}
}

func (inst* p4879aa5e6a_data_PhoneDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4879aa5e6.PhoneDao)
	nop(ie, com)

	


    return nil
}


