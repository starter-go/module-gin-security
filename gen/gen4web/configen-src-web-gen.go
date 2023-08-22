package gen4web
import (
    pd1a916a20 "github.com/starter-go/libgin"
    pa9413a197 "github.com/starter-go/module-security-gin-gorm/internal/web"
    pe1f113ef5 "github.com/starter-go/module-security-gin-gorm/internal/web/controllers"
    pa70cfe71f "github.com/starter-go/module-security-gin-gorm/internal/web/controllers/admin"
    pfbdcfa528 "github.com/starter-go/module-security-gin-gorm/internal/web/groups"
    p91f218d46 "github.com/starter-go/security/jwt"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type pa9413a197.GinJWTokenAdapter in package:github.com/starter-go/module-security-gin-gorm/internal/web
//
// id:com-a9413a1971fdc9d7-web-GinJWTokenAdapter
// class:class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type pa9413a1971_web_GinJWTokenAdapter struct {
}

func (inst* pa9413a1971_web_GinJWTokenAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a9413a1971fdc9d7-web-GinJWTokenAdapter"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa9413a1971_web_GinJWTokenAdapter) new() any {
    return &pa9413a197.GinJWTokenAdapter{}
}

func (inst* pa9413a1971_web_GinJWTokenAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa9413a197.GinJWTokenAdapter)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)
    com.CookieMaxAge = inst.getCookieMaxAge(ie)
    com.UseCookie = inst.getUseCookie(ie)
    com.UseHeader = inst.getUseHeader(ie)


    return nil
}


func (inst*pa9413a1971_web_GinJWTokenAdapter) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*pa9413a1971_web_GinJWTokenAdapter) getCookieMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${jwt.gin-adapter.cookie-max-age}")
}


func (inst*pa9413a1971_web_GinJWTokenAdapter) getUseCookie(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt.gin-adapter.use-cookie}")
}


func (inst*pa9413a1971_web_GinJWTokenAdapter) getUseHeader(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt.gin-adapter.use-header}")
}



// type pa9413a197.JSONGinResponder in package:github.com/starter-go/module-security-gin-gorm/internal/web
//
// id:com-a9413a1971fdc9d7-web-JSONGinResponder
// class:class-d1a916a203352fd5d33eabc36896b42e-Responder
// alias:
// scope:singleton
//
type pa9413a1971_web_JSONGinResponder struct {
}

func (inst* pa9413a1971_web_JSONGinResponder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a9413a1971fdc9d7-web-JSONGinResponder"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Responder"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa9413a1971_web_JSONGinResponder) new() any {
    return &pa9413a197.JSONGinResponder{}
}

func (inst* pa9413a1971_web_JSONGinResponder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa9413a197.JSONGinResponder)
	nop(ie, com)

	


    return nil
}



// type pe1f113ef5.AuthController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_AuthController struct {
}

func (inst* pe1f113ef5f_controllers_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_AuthController) new() any {
    return &pe1f113ef5.AuthController{}
}

func (inst* pe1f113ef5f_controllers_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.AuthController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.AuthSer = inst.getAuthSer(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_AuthController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_AuthController) getService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}


func (inst*pe1f113ef5f_controllers_AuthController) getAuthSer(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type pe1f113ef5.ExampleController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_ExampleController struct {
}

func (inst* pe1f113ef5f_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_ExampleController) new() any {
    return &pe1f113ef5.ExampleController{}
}

func (inst* pe1f113ef5f_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pe1f113ef5.PermissionController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_PermissionController struct {
}

func (inst* pe1f113ef5f_controllers_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_PermissionController) new() any {
    return &pe1f113ef5.PermissionController{}
}

func (inst* pe1f113ef5f_controllers_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pe1f113ef5.RoleController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_RoleController struct {
}

func (inst* pe1f113ef5f_controllers_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_RoleController) new() any {
    return &pe1f113ef5.RoleController{}
}

func (inst* pe1f113ef5f_controllers_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type pe1f113ef5.SessionController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_SessionController struct {
}

func (inst* pe1f113ef5f_controllers_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_SessionController) new() any {
    return &pe1f113ef5.SessionController{}
}

func (inst* pe1f113ef5f_controllers_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.SessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_SessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_SessionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*pe1f113ef5f_controllers_SessionController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type pe1f113ef5.TokenController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-TokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_TokenController struct {
}

func (inst* pe1f113ef5f_controllers_TokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-TokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_TokenController) new() any {
    return &pe1f113ef5.TokenController{}
}

func (inst* pe1f113ef5f_controllers_TokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.TokenController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_TokenController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_TokenController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*pe1f113ef5f_controllers_TokenController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type pe1f113ef5.UserController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers
//
// id:com-e1f113ef5fa1419f-controllers-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe1f113ef5f_controllers_UserController struct {
}

func (inst* pe1f113ef5f_controllers_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e1f113ef5fa1419f-controllers-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe1f113ef5f_controllers_UserController) new() any {
    return &pe1f113ef5.UserController{}
}

func (inst* pe1f113ef5f_controllers_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe1f113ef5.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pe1f113ef5f_controllers_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pe1f113ef5f_controllers_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}



// type pa70cfe71f.PermissionController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers/admin
//
// id:com-a70cfe71fbcbd740-admin-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pa70cfe71fb_admin_PermissionController struct {
}

func (inst* pa70cfe71fb_admin_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a70cfe71fbcbd740-admin-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa70cfe71fb_admin_PermissionController) new() any {
    return &pa70cfe71f.PermissionController{}
}

func (inst* pa70cfe71fb_admin_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa70cfe71f.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pa70cfe71fb_admin_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pa70cfe71fb_admin_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pa70cfe71f.RoleController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers/admin
//
// id:com-a70cfe71fbcbd740-admin-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pa70cfe71fb_admin_RoleController struct {
}

func (inst* pa70cfe71fb_admin_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a70cfe71fbcbd740-admin-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa70cfe71fb_admin_RoleController) new() any {
    return &pa70cfe71f.RoleController{}
}

func (inst* pa70cfe71fb_admin_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa70cfe71f.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pa70cfe71fb_admin_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pa70cfe71fb_admin_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type pa70cfe71f.UserController in package:github.com/starter-go/module-security-gin-gorm/internal/web/controllers/admin
//
// id:com-a70cfe71fbcbd740-admin-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pa70cfe71fb_admin_UserController struct {
}

func (inst* pa70cfe71fb_admin_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a70cfe71fbcbd740-admin-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa70cfe71fb_admin_UserController) new() any {
    return &pa70cfe71f.UserController{}
}

func (inst* pa70cfe71fb_admin_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa70cfe71f.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pa70cfe71fb_admin_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pa70cfe71fb_admin_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}



// type pfbdcfa528.AdminWebGroup in package:github.com/starter-go/module-security-gin-gorm/internal/web/groups
//
// id:com-fbdcfa5284ffd46f-groups-AdminWebGroup
// class:class-d1a916a203352fd5d33eabc36896b42e-Group
// alias:
// scope:singleton
//
type pfbdcfa5284_groups_AdminWebGroup struct {
}

func (inst* pfbdcfa5284_groups_AdminWebGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fbdcfa5284ffd46f-groups-AdminWebGroup"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Group"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfbdcfa5284_groups_AdminWebGroup) new() any {
    return &pfbdcfa528.AdminWebGroup{}
}

func (inst* pfbdcfa5284_groups_AdminWebGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfbdcfa528.AdminWebGroup)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Name = inst.getName(ie)
    com.Path = inst.getPath(ie)


    return nil
}


func (inst*pfbdcfa5284_groups_AdminWebGroup) getContext(ie application.InjectionExt)pd1a916a20.Context{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Context").(pd1a916a20.Context)
}


func (inst*pfbdcfa5284_groups_AdminWebGroup) getName(ie application.InjectionExt)string{
    return ie.GetString("${web-group.admin.name}")
}


func (inst*pfbdcfa5284_groups_AdminWebGroup) getPath(ie application.InjectionExt)string{
    return ie.GetString("${web-group.admin.path}")
}


