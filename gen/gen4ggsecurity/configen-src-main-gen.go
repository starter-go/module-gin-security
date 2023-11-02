package gen4ggsecurity
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p6a34f6f22 "github.com/starter-go/module-email/mails"
    p4f3ce922c "github.com/starter-go/security-gin-gorm/components/auth/auth1/email"
    p492cd0c9e "github.com/starter-go/security-gin-gorm/components/auth/auth1/sms"
    pc0223a2c4 "github.com/starter-go/security-gin-gorm/components/auth/auth1/vericode"
    p89ee959c3 "github.com/starter-go/security-gin-gorm/components/auth/auth2/login"
    p04a90047b "github.com/starter-go/security-gin-gorm/components/auth/auth2/signup"
    p56e4dad45 "github.com/starter-go/security-gin-gorm/components/services"
    p985949a7f "github.com/starter-go/security-gin-gorm/components/web/controllers/admin"
    p775fb156f "github.com/starter-go/security-gin-gorm/components/web/controllers/develop"
    p16ea5b788 "github.com/starter-go/security-gin-gorm/components/web/controllers/home"
    pf5d2c6fae "github.com/starter-go/security-gorm/rbacdb"
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type p4f3ce922c.AuthByEmail in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/email
//
// id:com-4f3ce922c52a0f87-email-AuthByEmail
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p4f3ce922c5_email_AuthByEmail struct {
}

func (inst* p4f3ce922c5_email_AuthByEmail) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4f3ce922c52a0f87-email-AuthByEmail"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4f3ce922c5_email_AuthByEmail) new() any {
    return &p4f3ce922c.AuthByEmail{}
}

func (inst* p4f3ce922c5_email_AuthByEmail) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4f3ce922c.AuthByEmail)
	nop(ie, com)

	
    com.VerificationService = inst.getVerificationService(ie)


    return nil
}


func (inst*p4f3ce922c5_email_AuthByEmail) getVerificationService(ie application.InjectionExt)p56e4dad45.VerificationService{
    return ie.GetComponent("#alias-56e4dad4531af2d8f160595779b3dfb6-VerificationService").(p56e4dad45.VerificationService)
}



// type p492cd0c9e.AuthBySMS in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/sms
//
// id:com-492cd0c9ee9c4374-sms-AuthBySMS
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p492cd0c9ee_sms_AuthBySMS struct {
}

func (inst* p492cd0c9ee_sms_AuthBySMS) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-492cd0c9ee9c4374-sms-AuthBySMS"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p492cd0c9ee_sms_AuthBySMS) new() any {
    return &p492cd0c9e.AuthBySMS{}
}

func (inst* p492cd0c9ee_sms_AuthBySMS) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p492cd0c9e.AuthBySMS)
	nop(ie, com)

	
    com.VerificationService = inst.getVerificationService(ie)


    return nil
}


func (inst*p492cd0c9ee_sms_AuthBySMS) getVerificationService(ie application.InjectionExt)p56e4dad45.VerificationService{
    return ie.GetComponent("#alias-56e4dad4531af2d8f160595779b3dfb6-VerificationService").(p56e4dad45.VerificationService)
}



// type pc0223a2c4.VerificationServiceImpl in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/vericode
//
// id:com-c0223a2c4a84ef9f-vericode-VerificationServiceImpl
// class:
// alias:alias-56e4dad4531af2d8f160595779b3dfb6-VerificationService
// scope:singleton
//
type pc0223a2c4a_vericode_VerificationServiceImpl struct {
}

func (inst* pc0223a2c4a_vericode_VerificationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c0223a2c4a84ef9f-vericode-VerificationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-56e4dad4531af2d8f160595779b3dfb6-VerificationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc0223a2c4a_vericode_VerificationServiceImpl) new() any {
    return &pc0223a2c4.VerificationServiceImpl{}
}

func (inst* pc0223a2c4a_vericode_VerificationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc0223a2c4.VerificationServiceImpl)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.JWT = inst.getJWT(ie)
    com.UUIDService = inst.getUUIDService(ie)
    com.SenderFromSMS = inst.getSenderFromSMS(ie)
    com.SenderFromMail = inst.getSenderFromMail(ie)
    com.TokenMaxAgeInMS = inst.getTokenMaxAgeInMS(ie)


    return nil
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getSender(ie application.InjectionExt)p6a34f6f22.Service{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-Service").(p6a34f6f22.Service)
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getJWT(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getSenderFromSMS(ie application.InjectionExt)string{
    return ie.GetString("${security.verification-code-sender.sms.from}")
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getSenderFromMail(ie application.InjectionExt)string{
    return ie.GetString("${security.verification-code-sender.mail.from}")
}


func (inst*pc0223a2c4a_vericode_VerificationServiceImpl) getTokenMaxAgeInMS(ie application.InjectionExt)int64{
    return ie.GetInt64("${security.verification-code-token.max-age.ms}")
}



// type p89ee959c3.AuthorizerForLogin in package:github.com/starter-go/security-gin-gorm/components/auth/auth2/login
//
// id:com-89ee959c3d3aa1e2-login-AuthorizerForLogin
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p89ee959c3d_login_AuthorizerForLogin struct {
}

func (inst* p89ee959c3d_login_AuthorizerForLogin) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-89ee959c3d3aa1e2-login-AuthorizerForLogin"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p89ee959c3d_login_AuthorizerForLogin) new() any {
    return &p89ee959c3.AuthorizerForLogin{}
}

func (inst* p89ee959c3d_login_AuthorizerForLogin) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p89ee959c3.AuthorizerForLogin)
	nop(ie, com)

	
    com.JWT = inst.getJWT(ie)


    return nil
}


func (inst*p89ee959c3d_login_AuthorizerForLogin) getJWT(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p04a90047b.AuthorizerForSignUp in package:github.com/starter-go/security-gin-gorm/components/auth/auth2/signup
//
// id:com-04a90047b85f3355-signup-AuthorizerForSignUp
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p04a90047b8_signup_AuthorizerForSignUp struct {
}

func (inst* p04a90047b8_signup_AuthorizerForSignUp) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-04a90047b85f3355-signup-AuthorizerForSignUp"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p04a90047b8_signup_AuthorizerForSignUp) new() any {
    return &p04a90047b.AuthorizerForSignUp{}
}

func (inst* p04a90047b8_signup_AuthorizerForSignUp) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p04a90047b.AuthorizerForSignUp)
	nop(ie, com)

	
    com.UserDAO = inst.getUserDAO(ie)
    com.PhoneDAO = inst.getPhoneDAO(ie)
    com.EmailDAO = inst.getEmailDAO(ie)
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p04a90047b8_signup_AuthorizerForSignUp) getUserDAO(ie application.InjectionExt)pf5d2c6fae.UserDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO").(pf5d2c6fae.UserDAO)
}


func (inst*p04a90047b8_signup_AuthorizerForSignUp) getPhoneDAO(ie application.InjectionExt)pf5d2c6fae.PhoneNumberDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-PhoneNumberDAO").(pf5d2c6fae.PhoneNumberDAO)
}


func (inst*p04a90047b8_signup_AuthorizerForSignUp) getEmailDAO(ie application.InjectionExt)pf5d2c6fae.EmailAddressDAO{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO").(pf5d2c6fae.EmailAddressDAO)
}


func (inst*p04a90047b8_signup_AuthorizerForSignUp) getAgent(ie application.InjectionExt)pf5d2c6fae.LocalAgent{
    return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}



// type p985949a7f.PermissionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-985949a7f73396fb-admin-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p985949a7f7_admin_PermissionController struct {
}

func (inst* p985949a7f7_admin_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-985949a7f73396fb-admin-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p985949a7f7_admin_PermissionController) new() any {
    return &p985949a7f.PermissionController{}
}

func (inst* p985949a7f7_admin_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p985949a7f.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p985949a7f7_admin_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p985949a7f7_admin_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p985949a7f.RegionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-985949a7f73396fb-admin-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p985949a7f7_admin_RegionController struct {
}

func (inst* p985949a7f7_admin_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-985949a7f73396fb-admin-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p985949a7f7_admin_RegionController) new() any {
    return &p985949a7f.RegionController{}
}

func (inst* p985949a7f7_admin_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p985949a7f.RegionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p985949a7f7_admin_RegionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p985949a7f7_admin_RegionController) getService(ie application.InjectionExt)p2dece1e49.RegionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}



// type p985949a7f.RoleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-985949a7f73396fb-admin-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p985949a7f7_admin_RoleController struct {
}

func (inst* p985949a7f7_admin_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-985949a7f73396fb-admin-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p985949a7f7_admin_RoleController) new() any {
    return &p985949a7f.RoleController{}
}

func (inst* p985949a7f7_admin_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p985949a7f.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p985949a7f7_admin_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p985949a7f7_admin_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type p985949a7f.UserController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-985949a7f73396fb-admin-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p985949a7f7_admin_UserController struct {
}

func (inst* p985949a7f7_admin_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-985949a7f73396fb-admin-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p985949a7f7_admin_UserController) new() any {
    return &p985949a7f.UserController{}
}

func (inst* p985949a7f7_admin_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p985949a7f.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p985949a7f7_admin_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p985949a7f7_admin_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}



// type p775fb156f.DebugController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/develop
//
// id:com-775fb156f4ed4e07-develop-DebugController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p775fb156f4_develop_DebugController struct {
}

func (inst* p775fb156f4_develop_DebugController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-775fb156f4ed4e07-develop-DebugController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p775fb156f4_develop_DebugController) new() any {
    return &p775fb156f.DebugController{}
}

func (inst* p775fb156f4_develop_DebugController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p775fb156f.DebugController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p775fb156f4_develop_DebugController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p775fb156f4_develop_DebugController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p775fb156f.ExampleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/develop
//
// id:com-775fb156f4ed4e07-develop-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p775fb156f4_develop_ExampleController struct {
}

func (inst* p775fb156f4_develop_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-775fb156f4ed4e07-develop-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p775fb156f4_develop_ExampleController) new() any {
    return &p775fb156f.ExampleController{}
}

func (inst* p775fb156f4_develop_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p775fb156f.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p775fb156f4_develop_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p775fb156f4_develop_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p16ea5b788.AuthController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_AuthController struct {
}

func (inst* p16ea5b788f_home_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_AuthController) new() any {
    return &p16ea5b788.AuthController{}
}

func (inst* p16ea5b788f_home_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.AuthController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.AuthSer = inst.getAuthSer(ie)


    return nil
}


func (inst*p16ea5b788f_home_AuthController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_AuthController) getService(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}


func (inst*p16ea5b788f_home_AuthController) getAuthSer(ie application.InjectionExt)p2dece1e49.AuthService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}



// type p16ea5b788.ExampleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_ExampleController struct {
}

func (inst* p16ea5b788f_home_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_ExampleController) new() any {
    return &p16ea5b788.ExampleController{}
}

func (inst* p16ea5b788f_home_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p16ea5b788f_home_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_ExampleController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p16ea5b788.PermissionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_PermissionController struct {
}

func (inst* p16ea5b788f_home_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_PermissionController) new() any {
    return &p16ea5b788.PermissionController{}
}

func (inst* p16ea5b788f_home_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.PermissionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p16ea5b788f_home_PermissionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_PermissionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}



// type p16ea5b788.RegionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_RegionController struct {
}

func (inst* p16ea5b788f_home_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_RegionController) new() any {
    return &p16ea5b788.RegionController{}
}

func (inst* p16ea5b788f_home_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.RegionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p16ea5b788f_home_RegionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_RegionController) getService(ie application.InjectionExt)p2dece1e49.RegionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}



// type p16ea5b788.RoleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_RoleController struct {
}

func (inst* p16ea5b788f_home_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_RoleController) new() any {
    return &p16ea5b788.RoleController{}
}

func (inst* p16ea5b788f_home_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.RoleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p16ea5b788f_home_RoleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_RoleController) getService(ie application.InjectionExt)p2dece1e49.RoleService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}



// type p16ea5b788.SessionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_SessionController struct {
}

func (inst* p16ea5b788f_home_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_SessionController) new() any {
    return &p16ea5b788.SessionController{}
}

func (inst* p16ea5b788f_home_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.SessionController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*p16ea5b788f_home_SessionController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_SessionController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*p16ea5b788f_home_SessionController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p16ea5b788.TokenController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-TokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_TokenController struct {
}

func (inst* p16ea5b788f_home_TokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-TokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_TokenController) new() any {
    return &p16ea5b788.TokenController{}
}

func (inst* p16ea5b788f_home_TokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.TokenController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)
    com.JWTService = inst.getJWTService(ie)


    return nil
}


func (inst*p16ea5b788f_home_TokenController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_TokenController) getService(ie application.InjectionExt)p2dece1e49.PermissionService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}


func (inst*p16ea5b788f_home_TokenController) getJWTService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p16ea5b788.UserController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-16ea5b788fa3dca4-home-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p16ea5b788f_home_UserController struct {
}

func (inst* p16ea5b788f_home_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-16ea5b788fa3dca4-home-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p16ea5b788f_home_UserController) new() any {
    return &p16ea5b788.UserController{}
}

func (inst* p16ea5b788f_home_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p16ea5b788.UserController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p16ea5b788f_home_UserController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p16ea5b788f_home_UserController) getService(ie application.InjectionExt)p2dece1e49.UserService{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}


