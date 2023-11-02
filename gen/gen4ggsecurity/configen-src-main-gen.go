package gen4ggsecurity

import (
	"github.com/starter-go/application"
	pd1a916a20 "github.com/starter-go/libgin"
	p6a34f6f22 "github.com/starter-go/module-email/mails"
	pf48f932fb "github.com/starter-go/security-gin-gorm/components/auth/auth1/email"
	pd4be75b3b "github.com/starter-go/security-gin-gorm/components/auth/auth1/sms"
	pbd5217dfa "github.com/starter-go/security-gin-gorm/components/auth/auth1/vericode"
	pdc51ff5e4 "github.com/starter-go/security-gin-gorm/components/auth/auth2/login"
	p7f9a7d1cf "github.com/starter-go/security-gin-gorm/components/auth/auth2/signup"
	p8e41bd317 "github.com/starter-go/security-gin-gorm/components/services"
	pef5fb8e15 "github.com/starter-go/security-gin-gorm/components/web/controllers/admin"
	p13ed4c221 "github.com/starter-go/security-gin-gorm/components/web/controllers/develop"
	p736bdaa37 "github.com/starter-go/security-gin-gorm/components/web/controllers/home"
	pf5d2c6fae "github.com/starter-go/security-gorm/rbacdb"
	p91f218d46 "github.com/starter-go/security/jwt"
	p9621e8b71 "github.com/starter-go/security/random"
	p2dece1e49 "github.com/starter-go/security/rbac"
)

// type pf48f932fb.AuthByEmail in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/email
//
// id:com-f48f932fb4a915f5-email-AuthByEmail
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type pf48f932fb4_email_AuthByEmail struct {
}

func (inst *pf48f932fb4_email_AuthByEmail) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f48f932fb4a915f5-email-AuthByEmail"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pf48f932fb4_email_AuthByEmail) new() any {
	return &pf48f932fb.AuthByEmail{}
}

func (inst *pf48f932fb4_email_AuthByEmail) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf48f932fb.AuthByEmail)
	nop(ie, com)

	com.VerificationService = inst.getVerificationService(ie)

	return nil
}

func (inst *pf48f932fb4_email_AuthByEmail) getVerificationService(ie application.InjectionExt) p8e41bd317.VerificationService {
	return ie.GetComponent("#alias-8e41bd317406e4801e90c494ff828ee2-VerificationService").(p8e41bd317.VerificationService)
}

// type pd4be75b3b.AuthBySMS in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/sms
//
// id:com-d4be75b3b4baaa69-sms-AuthBySMS
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type pd4be75b3b4_sms_AuthBySMS struct {
}

func (inst *pd4be75b3b4_sms_AuthBySMS) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d4be75b3b4baaa69-sms-AuthBySMS"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pd4be75b3b4_sms_AuthBySMS) new() any {
	return &pd4be75b3b.AuthBySMS{}
}

func (inst *pd4be75b3b4_sms_AuthBySMS) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd4be75b3b.AuthBySMS)
	nop(ie, com)

	com.VerificationService = inst.getVerificationService(ie)

	return nil
}

func (inst *pd4be75b3b4_sms_AuthBySMS) getVerificationService(ie application.InjectionExt) p8e41bd317.VerificationService {
	return ie.GetComponent("#alias-8e41bd317406e4801e90c494ff828ee2-VerificationService").(p8e41bd317.VerificationService)
}

// type pbd5217dfa.VerificationServiceImpl in package:github.com/starter-go/security-gin-gorm/components/auth/auth1/vericode
//
// id:com-bd5217dfaec2d26c-vericode-VerificationServiceImpl
// class:
// alias:alias-8e41bd317406e4801e90c494ff828ee2-VerificationService
// scope:singleton
//
type pbd5217dfae_vericode_VerificationServiceImpl struct {
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bd5217dfaec2d26c-vericode-VerificationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-8e41bd317406e4801e90c494ff828ee2-VerificationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) new() any {
	return &pbd5217dfa.VerificationServiceImpl{}
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbd5217dfa.VerificationServiceImpl)
	nop(ie, com)

	com.Sender = inst.getSender(ie)
	com.JWT = inst.getJWT(ie)
	com.UUIDService = inst.getUUIDService(ie)
	com.SenderFromSMS = inst.getSenderFromSMS(ie)
	com.SenderFromMail = inst.getSenderFromMail(ie)
	com.TokenMaxAgeInMS = inst.getTokenMaxAgeInMS(ie)

	return nil
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getSender(ie application.InjectionExt) p6a34f6f22.Service {
	return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-Service").(p6a34f6f22.Service)
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getJWT(ie application.InjectionExt) p91f218d46.Service {
	return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getUUIDService(ie application.InjectionExt) p9621e8b71.UUIDService {
	return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getSenderFromSMS(ie application.InjectionExt) string {
	return ie.GetString("${security.verification-code-sender.sms.from}")
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getSenderFromMail(ie application.InjectionExt) string {
	return ie.GetString("${security.verification-code-sender.mail.from}")
}

func (inst *pbd5217dfae_vericode_VerificationServiceImpl) getTokenMaxAgeInMS(ie application.InjectionExt) int64 {
	return ie.GetInt64("${security.verification-code-token.max-age.ms}")
}

// type pdc51ff5e4.AuthorizerForLogin in package:github.com/starter-go/security-gin-gorm/components/auth/auth2/login
//
// id:com-dc51ff5e4d3745d8-login-AuthorizerForLogin
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type pdc51ff5e4d_login_AuthorizerForLogin struct {
}

func (inst *pdc51ff5e4d_login_AuthorizerForLogin) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dc51ff5e4d3745d8-login-AuthorizerForLogin"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pdc51ff5e4d_login_AuthorizerForLogin) new() any {
	return &pdc51ff5e4.AuthorizerForLogin{}
}

func (inst *pdc51ff5e4d_login_AuthorizerForLogin) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdc51ff5e4.AuthorizerForLogin)
	nop(ie, com)

	com.JWT = inst.getJWT(ie)

	return nil
}

func (inst *pdc51ff5e4d_login_AuthorizerForLogin) getJWT(ie application.InjectionExt) p91f218d46.Service {
	return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}

// type p7f9a7d1cf.AuthorizerForSignUp in package:github.com/starter-go/security-gin-gorm/components/auth/auth2/signup
//
// id:com-7f9a7d1cfc2f8e9e-signup-AuthorizerForSignUp
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p7f9a7d1cfc_signup_AuthorizerForSignUp struct {
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7f9a7d1cfc2f8e9e-signup-AuthorizerForSignUp"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) new() any {
	return &p7f9a7d1cf.AuthorizerForSignUp{}
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7f9a7d1cf.AuthorizerForSignUp)
	nop(ie, com)

	com.UserDAO = inst.getUserDAO(ie)
	com.PhoneDAO = inst.getPhoneDAO(ie)
	com.EmailDAO = inst.getEmailDAO(ie)
	com.Agent = inst.getAgent(ie)

	return nil
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) getUserDAO(ie application.InjectionExt) pf5d2c6fae.UserDAO {
	return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-UserDAO").(pf5d2c6fae.UserDAO)
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) getPhoneDAO(ie application.InjectionExt) pf5d2c6fae.PhoneNumberDAO {
	return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-PhoneNumberDAO").(pf5d2c6fae.PhoneNumberDAO)
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) getEmailDAO(ie application.InjectionExt) pf5d2c6fae.EmailAddressDAO {
	return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-EmailAddressDAO").(pf5d2c6fae.EmailAddressDAO)
}

func (inst *p7f9a7d1cfc_signup_AuthorizerForSignUp) getAgent(ie application.InjectionExt) pf5d2c6fae.LocalAgent {
	return ie.GetComponent("#alias-f5d2c6fae036814399fa2ed06c0dc99f-LocalAgent").(pf5d2c6fae.LocalAgent)
}

// type pef5fb8e15.PermissionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_PermissionController struct {
}

func (inst *pef5fb8e156_admin_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pef5fb8e156_admin_PermissionController) new() any {
	return &pef5fb8e15.PermissionController{}
}

func (inst *pef5fb8e156_admin_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.PermissionController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *pef5fb8e156_admin_PermissionController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *pef5fb8e156_admin_PermissionController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

// type pef5fb8e15.RegionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_RegionController struct {
}

func (inst *pef5fb8e156_admin_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pef5fb8e156_admin_RegionController) new() any {
	return &pef5fb8e15.RegionController{}
}

func (inst *pef5fb8e156_admin_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.RegionController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *pef5fb8e156_admin_RegionController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *pef5fb8e156_admin_RegionController) getService(ie application.InjectionExt) p2dece1e49.RegionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}

// type pef5fb8e15.RoleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_RoleController struct {
}

func (inst *pef5fb8e156_admin_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pef5fb8e156_admin_RoleController) new() any {
	return &pef5fb8e15.RoleController{}
}

func (inst *pef5fb8e156_admin_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.RoleController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *pef5fb8e156_admin_RoleController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *pef5fb8e156_admin_RoleController) getService(ie application.InjectionExt) p2dece1e49.RoleService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}

// type pef5fb8e15.UserController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/admin
//
// id:com-ef5fb8e1568f4328-admin-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pef5fb8e156_admin_UserController struct {
}

func (inst *pef5fb8e156_admin_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef5fb8e1568f4328-admin-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pef5fb8e156_admin_UserController) new() any {
	return &pef5fb8e15.UserController{}
}

func (inst *pef5fb8e156_admin_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef5fb8e15.UserController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *pef5fb8e156_admin_UserController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *pef5fb8e156_admin_UserController) getService(ie application.InjectionExt) p2dece1e49.UserService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}

// type p13ed4c221.DebugController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/develop
//
// id:com-13ed4c221a66afef-develop-DebugController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p13ed4c221a_develop_DebugController struct {
}

func (inst *p13ed4c221a_develop_DebugController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-13ed4c221a66afef-develop-DebugController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p13ed4c221a_develop_DebugController) new() any {
	return &p13ed4c221.DebugController{}
}

func (inst *p13ed4c221a_develop_DebugController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p13ed4c221.DebugController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p13ed4c221a_develop_DebugController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p13ed4c221a_develop_DebugController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

// type p13ed4c221.ExampleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/develop
//
// id:com-13ed4c221a66afef-develop-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p13ed4c221a_develop_ExampleController struct {
}

func (inst *p13ed4c221a_develop_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-13ed4c221a66afef-develop-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p13ed4c221a_develop_ExampleController) new() any {
	return &p13ed4c221.ExampleController{}
}

func (inst *p13ed4c221a_develop_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p13ed4c221.ExampleController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p13ed4c221a_develop_ExampleController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p13ed4c221a_develop_ExampleController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

// type p736bdaa37.AuthController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_AuthController struct {
}

func (inst *p736bdaa375_home_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_AuthController) new() any {
	return &p736bdaa37.AuthController{}
}

func (inst *p736bdaa375_home_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.AuthController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)
	com.AuthSer = inst.getAuthSer(ie)

	return nil
}

func (inst *p736bdaa375_home_AuthController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_AuthController) getService(ie application.InjectionExt) p2dece1e49.AuthService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}

func (inst *p736bdaa375_home_AuthController) getAuthSer(ie application.InjectionExt) p2dece1e49.AuthService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-AuthService").(p2dece1e49.AuthService)
}

// type p736bdaa37.ExampleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_ExampleController struct {
}

func (inst *p736bdaa375_home_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_ExampleController) new() any {
	return &p736bdaa37.ExampleController{}
}

func (inst *p736bdaa375_home_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.ExampleController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p736bdaa375_home_ExampleController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_ExampleController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

// type p736bdaa37.PermissionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-PermissionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_PermissionController struct {
}

func (inst *p736bdaa375_home_PermissionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-PermissionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_PermissionController) new() any {
	return &p736bdaa37.PermissionController{}
}

func (inst *p736bdaa375_home_PermissionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.PermissionController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p736bdaa375_home_PermissionController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_PermissionController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

// type p736bdaa37.RegionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-RegionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_RegionController struct {
}

func (inst *p736bdaa375_home_RegionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-RegionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_RegionController) new() any {
	return &p736bdaa37.RegionController{}
}

func (inst *p736bdaa375_home_RegionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.RegionController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p736bdaa375_home_RegionController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_RegionController) getService(ie application.InjectionExt) p2dece1e49.RegionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RegionService").(p2dece1e49.RegionService)
}

// type p736bdaa37.RoleController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-RoleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_RoleController struct {
}

func (inst *p736bdaa375_home_RoleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-RoleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_RoleController) new() any {
	return &p736bdaa37.RoleController{}
}

func (inst *p736bdaa375_home_RoleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.RoleController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p736bdaa375_home_RoleController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_RoleController) getService(ie application.InjectionExt) p2dece1e49.RoleService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleService").(p2dece1e49.RoleService)
}

// type p736bdaa37.SessionController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-SessionController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_SessionController struct {
}

func (inst *p736bdaa375_home_SessionController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-SessionController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_SessionController) new() any {
	return &p736bdaa37.SessionController{}
}

func (inst *p736bdaa375_home_SessionController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.SessionController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)
	com.JWTService = inst.getJWTService(ie)

	return nil
}

func (inst *p736bdaa375_home_SessionController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_SessionController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

func (inst *p736bdaa375_home_SessionController) getJWTService(ie application.InjectionExt) p91f218d46.Service {
	return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}

// type p736bdaa37.TokenController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-TokenController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_TokenController struct {
}

func (inst *p736bdaa375_home_TokenController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-TokenController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_TokenController) new() any {
	return &p736bdaa37.TokenController{}
}

func (inst *p736bdaa375_home_TokenController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.TokenController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)
	com.JWTService = inst.getJWTService(ie)

	return nil
}

func (inst *p736bdaa375_home_TokenController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_TokenController) getService(ie application.InjectionExt) p2dece1e49.PermissionService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionService").(p2dece1e49.PermissionService)
}

func (inst *p736bdaa375_home_TokenController) getJWTService(ie application.InjectionExt) p91f218d46.Service {
	return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}

// type p736bdaa37.UserController in package:github.com/starter-go/security-gin-gorm/components/web/controllers/home
//
// id:com-736bdaa37581c9dc-home-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p736bdaa375_home_UserController struct {
}

func (inst *p736bdaa375_home_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-736bdaa37581c9dc-home-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *p736bdaa375_home_UserController) new() any {
	return &p736bdaa37.UserController{}
}

func (inst *p736bdaa375_home_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p736bdaa37.UserController)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.Service = inst.getService(ie)

	return nil
}

func (inst *p736bdaa375_home_UserController) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *p736bdaa375_home_UserController) getService(ie application.InjectionExt) p2dece1e49.UserService {
	return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserService").(p2dece1e49.UserService)
}
