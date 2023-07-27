package gen4main
import (
    pd1a916a20 "github.com/starter-go/libgin"
    ped246966f "github.com/starter-go/module-gin-security/src/main/golang"
    pda2eabb8f "github.com/starter-go/module-gin-security/src/main/golang/controllers"
    p4ba0ff4dd "github.com/starter-go/module-gin-security/src/main/golang/tokens"
    p91f218d46 "github.com/starter-go/security/jwt"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type ped246966f.JSONGinResponder in package:github.com/starter-go/module-gin-security/src/main/golang
//
// id:com-ed246966fe53fc24-golang-JSONGinResponder
// class:class-d1a916a203352fd5d33eabc36896b42e-Responder
// alias:
// scope:singleton
//
type ped246966fe_golang_JSONGinResponder struct {
}

func (inst* ped246966fe_golang_JSONGinResponder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ed246966fe53fc24-golang-JSONGinResponder"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Responder"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* ped246966fe_golang_JSONGinResponder) new() any {
    return &ped246966f.JSONGinResponder{}
}

func (inst* ped246966fe_golang_JSONGinResponder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*ped246966f.JSONGinResponder)
	nop(ie, com)

	


    return nil
}



// type pda2eabb8f.AuthController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_AuthController struct {
}

func (inst* pda2eabb8fd_controllers_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_AuthController) new() any {
    return &pda2eabb8f.AuthController{}
}

func (inst* pda2eabb8fd_controllers_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.AuthController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_AuthController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_AuthController) getService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type pda2eabb8f.ExampleController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_ExampleController struct {
}

func (inst* pda2eabb8fd_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_ExampleController) new() any {
    return &pda2eabb8f.ExampleController{}
}

func (inst* pda2eabb8fd_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pda2eabb8f.PermissionController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_PermissionController struct {
}

func (inst* pda2eabb8fd_controllers_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_PermissionController) new() any {
    return &pda2eabb8f.PermissionController{}
}

func (inst* pda2eabb8fd_controllers_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pda2eabb8f.RoleController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_RoleController struct {
}

func (inst* pda2eabb8fd_controllers_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_RoleController) new() any {
    return &pda2eabb8f.RoleController{}
}

func (inst* pda2eabb8fd_controllers_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type pda2eabb8f.SessionController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_SessionController struct {
}

func (inst* pda2eabb8fd_controllers_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_SessionController) new() any {
    return &pda2eabb8f.SessionController{}
}

func (inst* pda2eabb8fd_controllers_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.SessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_SessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_SessionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type pda2eabb8f.UserController in package:github.com/starter-go/module-gin-security/src/main/golang/controllers
//
// id:com-da2eabb8fdacb514-controllers-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pda2eabb8fd_controllers_UserController struct {
}

func (inst* pda2eabb8fd_controllers_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-da2eabb8fdacb514-controllers-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pda2eabb8fd_controllers_UserController) new() any {
    return &pda2eabb8f.UserController{}
}

func (inst* pda2eabb8fd_controllers_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pda2eabb8f.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pda2eabb8fd_controllers_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pda2eabb8fd_controllers_UserController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p4ba0ff4dd.GinJWTokenAdapter in package:github.com/starter-go/module-gin-security/src/main/golang/tokens
//
// id:com-4ba0ff4ddac0e0ef-tokens-GinJWTokenAdapter
// class:class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p4ba0ff4dda_tokens_GinJWTokenAdapter struct {
}

func (inst* p4ba0ff4dda_tokens_GinJWTokenAdapter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4ba0ff4ddac0e0ef-tokens-GinJWTokenAdapter"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Adapter class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4ba0ff4dda_tokens_GinJWTokenAdapter) new() any {
    return &p4ba0ff4dd.GinJWTokenAdapter{}
}

func (inst* p4ba0ff4dda_tokens_GinJWTokenAdapter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4ba0ff4dd.GinJWTokenAdapter)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)
    com.CookieMaxAge = inst.getCookieMaxAge(ie)
    com.UseCookie = inst.getUseCookie(ie)
    com.UseHeader = inst.getUseHeader(ie)


    return nil
}


func (inst*p4ba0ff4dda_tokens_GinJWTokenAdapter) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p4ba0ff4dda_tokens_GinJWTokenAdapter) getCookieMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${jwt.gin-adapter.cookie-max-age}")
}


func (inst*p4ba0ff4dda_tokens_GinJWTokenAdapter) getUseCookie(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt.gin-adapter.use-cookie}")
}


func (inst*p4ba0ff4dda_tokens_GinJWTokenAdapter) getUseHeader(ie application.InjectionExt)bool{
    return ie.GetBool("${jwt.gin-adapter.use-header}")
}


