package gen4ggsecurity
import (
    pd1a916a20 "github.com/starter-go/libgin"
    pef5fb8e15 "github.com/starter-go/module-security-gin-gorm/components/web/controllers/admin"
    p13ed4c221 "github.com/starter-go/module-security-gin-gorm/components/web/controllers/develop"
    p736bdaa37 "github.com/starter-go/module-security-gin-gorm/components/web/controllers/home"
    p91f218d46 "github.com/starter-go/security/jwt"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type pef5fb8e15.PermissionController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_PermissionController struct {
}

func (inst* pef5fb8e156_admin_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef5fb8e156_admin_PermissionController) new() any {
    return &pef5fb8e15.PermissionController{}
}

func (inst* pef5fb8e156_admin_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pef5fb8e156_admin_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pef5fb8e156_admin_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pef5fb8e15.RegionController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_RegionController struct {
}

func (inst* pef5fb8e156_admin_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef5fb8e156_admin_RegionController) new() any {
    return &pef5fb8e15.RegionController{}
}

func (inst* pef5fb8e156_admin_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.RegionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pef5fb8e156_admin_RegionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pef5fb8e156_admin_RegionController) getService(ie application.InjectionExt)p2dece1e49.RegionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}



// type pef5fb8e15.RoleController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_RoleController struct {
}

func (inst* pef5fb8e156_admin_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef5fb8e156_admin_RoleController) new() any {
    return &pef5fb8e15.RoleController{}
}

func (inst* pef5fb8e156_admin_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pef5fb8e156_admin_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pef5fb8e156_admin_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type pef5fb8e15.UserController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_UserController struct {
}

func (inst* pef5fb8e156_admin_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef5fb8e156_admin_UserController) new() any {
    return &pef5fb8e15.UserController{}
}

func (inst* pef5fb8e156_admin_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pef5fb8e156_admin_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pef5fb8e156_admin_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}



// type p13ed4c221.DebugController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/develop
//
// id:com-13ed4c221a66afef-develop-DebugController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p13ed4c221a_develop_DebugController struct {
}

func (inst* p13ed4c221a_develop_DebugController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-13ed4c221a66afef-develop-DebugController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p13ed4c221a_develop_DebugController) new() any {
    return &p13ed4c221.DebugController{}
}

func (inst* p13ed4c221a_develop_DebugController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p13ed4c221.DebugController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p13ed4c221a_develop_DebugController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p13ed4c221a_develop_DebugController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p13ed4c221.ExampleController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/develop
//
// id:com-13ed4c221a66afef-develop-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p13ed4c221a_develop_ExampleController struct {
}

func (inst* p13ed4c221a_develop_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-13ed4c221a66afef-develop-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p13ed4c221a_develop_ExampleController) new() any {
    return &p13ed4c221.ExampleController{}
}

func (inst* p13ed4c221a_develop_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p13ed4c221.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p13ed4c221a_develop_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p13ed4c221a_develop_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p736bdaa37.AuthController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_AuthController struct {
}

func (inst* p736bdaa375_home_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_AuthController) new() any {
    return &p736bdaa37.AuthController{}
}

func (inst* p736bdaa375_home_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.AuthController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.AuthSer = inst.getAuthSer(ie)


    return nil
}


func (inst*p736bdaa375_home_AuthController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_AuthController) getService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}


func (inst*p736bdaa375_home_AuthController) getAuthSer(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type p736bdaa37.ExampleController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_ExampleController struct {
}

func (inst* p736bdaa375_home_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_ExampleController) new() any {
    return &p736bdaa37.ExampleController{}
}

func (inst* p736bdaa375_home_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p736bdaa375_home_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p736bdaa37.PermissionController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_PermissionController struct {
}

func (inst* p736bdaa375_home_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_PermissionController) new() any {
    return &p736bdaa37.PermissionController{}
}

func (inst* p736bdaa375_home_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p736bdaa375_home_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p736bdaa37.RegionController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_RegionController struct {
}

func (inst* p736bdaa375_home_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_RegionController) new() any {
    return &p736bdaa37.RegionController{}
}

func (inst* p736bdaa375_home_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.RegionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p736bdaa375_home_RegionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_RegionController) getService(ie application.InjectionExt)p2dece1e49.RegionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}



// type p736bdaa37.RoleController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_RoleController struct {
}

func (inst* p736bdaa375_home_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_RoleController) new() any {
    return &p736bdaa37.RoleController{}
}

func (inst* p736bdaa375_home_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p736bdaa375_home_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type p736bdaa37.SessionController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_SessionController struct {
}

func (inst* p736bdaa375_home_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_SessionController) new() any {
    return &p736bdaa37.SessionController{}
}

func (inst* p736bdaa375_home_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.SessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*p736bdaa375_home_SessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_SessionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*p736bdaa375_home_SessionController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p736bdaa37.TokenController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-TokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_TokenController struct {
}

func (inst* p736bdaa375_home_TokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-TokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_TokenController) new() any {
    return &p736bdaa37.TokenController{}
}

func (inst* p736bdaa375_home_TokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.TokenController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*p736bdaa375_home_TokenController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_TokenController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*p736bdaa375_home_TokenController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p736bdaa37.UserController in package:github.com/starter-go/module-security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_UserController struct {
}

func (inst* p736bdaa375_home_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p736bdaa375_home_UserController) new() any {
    return &p736bdaa37.UserController{}
}

func (inst* p736bdaa375_home_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p736bdaa375_home_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p736bdaa375_home_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}


