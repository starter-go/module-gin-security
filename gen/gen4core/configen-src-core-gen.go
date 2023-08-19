package gen4core
import (
    p512a30914 "github.com/starter-go/libgorm"
    p6a34f6f22 "github.com/starter-go/module-email/mails"
    p623d365fa "github.com/starter-go/module-security-gin-gorm/internal/core/authx"
    pcb7df5f5e "github.com/starter-go/module-security-gin-gorm/internal/core/permissions"
    pd8c903054 "github.com/starter-go/module-security-gin-gorm/internal/core/roles"
    p12d48b32e "github.com/starter-go/module-security-gin-gorm/internal/core/sessions"
    p6bb98f719 "github.com/starter-go/module-security-gin-gorm/internal/core/tokens"
    p217db24b0 "github.com/starter-go/module-security-gin-gorm/internal/core/users"
    p9196ff9ab "github.com/starter-go/module-security-gin-gorm/internal/services"
    p91f218d46 "github.com/starter-go/security/jwt"
    p9621e8b71 "github.com/starter-go/security/random"
    p2dece1e49 "github.com/starter-go/security/rbac"
     "github.com/starter-go/application"
)

// type p623d365fa.AuthWithEmail in package:github.com/starter-go/module-security-gin-gorm/internal/core/authx
//
// id:com-623d365faaa9ccfd-authx-AuthWithEmail
// class:class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p623d365faa_authx_AuthWithEmail struct {
}

func (inst* p623d365faa_authx_AuthWithEmail) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-623d365faaa9ccfd-authx-AuthWithEmail"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p623d365faa_authx_AuthWithEmail) new() any {
    return &p623d365fa.AuthWithEmail{}
}

func (inst* p623d365faa_authx_AuthWithEmail) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p623d365fa.AuthWithEmail)
	nop(ie, com)

	
    com.RandomSer = inst.getRandomSer(ie)
    com.MailSer = inst.getMailSer(ie)
    com.JWTSer = inst.getJWTSer(ie)
    com.VeriSer = inst.getVeriSer(ie)
    com.AuthorSer = inst.getAuthorSer(ie)
    com.DBAgent = inst.getDBAgent(ie)
    com.Users = inst.getUsers(ie)
    com.EmailAddresses = inst.getEmailAddresses(ie)
    com.PhoneNumbers = inst.getPhoneNumbers(ie)


    return nil
}


func (inst*p623d365faa_authx_AuthWithEmail) getRandomSer(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*p623d365faa_authx_AuthWithEmail) getMailSer(ie application.InjectionExt)p6a34f6f22.Service{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-Service").(p6a34f6f22.Service)
}


func (inst*p623d365faa_authx_AuthWithEmail) getJWTSer(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p623d365faa_authx_AuthWithEmail) getVeriSer(ie application.InjectionExt)p9196ff9ab.VerificationService{
    return ie.GetComponent("#alias-9196ff9ab6d9e7bcc043594cacdbd432-VerificationService").(p9196ff9ab.VerificationService)
}


func (inst*p623d365faa_authx_AuthWithEmail) getAuthorSer(ie application.InjectionExt)p9196ff9ab.AuthorizationService{
    return ie.GetComponent("#alias-9196ff9ab6d9e7bcc043594cacdbd432-AuthorizationService").(p9196ff9ab.AuthorizationService)
}


func (inst*p623d365faa_authx_AuthWithEmail) getDBAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p623d365faa_authx_AuthWithEmail) getUsers(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}


func (inst*p623d365faa_authx_AuthWithEmail) getEmailAddresses(ie application.InjectionExt)p2dece1e49.EmailAddressDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-EmailAddressDAO").(p2dece1e49.EmailAddressDAO)
}


func (inst*p623d365faa_authx_AuthWithEmail) getPhoneNumbers(ie application.InjectionExt)p2dece1e49.PhoneNumberDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PhoneNumberDAO").(p2dece1e49.PhoneNumberDAO)
}



// type p623d365fa.PasswordAuth in package:github.com/starter-go/module-security-gin-gorm/internal/core/authx
//
// id:com-623d365faaa9ccfd-authx-PasswordAuth
// class:class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry
// alias:
// scope:singleton
//
type p623d365faa_authx_PasswordAuth struct {
}

func (inst* p623d365faa_authx_PasswordAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-623d365faaa9ccfd-authx-PasswordAuth"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Authenticator class-9d209f7c2504d33e6054a2c9998e9485-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p623d365faa_authx_PasswordAuth) new() any {
    return &p623d365fa.PasswordAuth{}
}

func (inst* p623d365faa_authx_PasswordAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p623d365fa.PasswordAuth)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)
    com.AuthSer = inst.getAuthSer(ie)


    return nil
}


func (inst*p623d365faa_authx_PasswordAuth) getUserDao(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}


func (inst*p623d365faa_authx_PasswordAuth) getAuthSer(ie application.InjectionExt)p9196ff9ab.AuthorizationService{
    return ie.GetComponent("#alias-9196ff9ab6d9e7bcc043594cacdbd432-AuthorizationService").(p9196ff9ab.AuthorizationService)
}



// type p623d365fa.AuthWithSMS in package:github.com/starter-go/module-security-gin-gorm/internal/core/authx
//
// id:com-623d365faaa9ccfd-authx-AuthWithSMS
// class:class-9d209f7c2504d33e6054a2c9998e9485-Authenticator
// alias:
// scope:singleton
//
type p623d365faa_authx_AuthWithSMS struct {
}

func (inst* p623d365faa_authx_AuthWithSMS) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-623d365faaa9ccfd-authx-AuthWithSMS"
	r.Classes = "class-9d209f7c2504d33e6054a2c9998e9485-Authenticator"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p623d365faa_authx_AuthWithSMS) new() any {
    return &p623d365fa.AuthWithSMS{}
}

func (inst* p623d365faa_authx_AuthWithSMS) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p623d365fa.AuthWithSMS)
	nop(ie, com)

	
    com.RandomSer = inst.getRandomSer(ie)
    com.MailSer = inst.getMailSer(ie)
    com.JWTSer = inst.getJWTSer(ie)
    com.VeriSer = inst.getVeriSer(ie)
    com.AuthorSer = inst.getAuthorSer(ie)
    com.DBAgent = inst.getDBAgent(ie)
    com.Users = inst.getUsers(ie)
    com.EmailAddresses = inst.getEmailAddresses(ie)
    com.PhoneNumbers = inst.getPhoneNumbers(ie)


    return nil
}


func (inst*p623d365faa_authx_AuthWithSMS) getRandomSer(ie application.InjectionExt)p9621e8b71.Service{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-Service").(p9621e8b71.Service)
}


func (inst*p623d365faa_authx_AuthWithSMS) getMailSer(ie application.InjectionExt)p6a34f6f22.Service{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-Service").(p6a34f6f22.Service)
}


func (inst*p623d365faa_authx_AuthWithSMS) getJWTSer(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p623d365faa_authx_AuthWithSMS) getVeriSer(ie application.InjectionExt)p9196ff9ab.VerificationService{
    return ie.GetComponent("#alias-9196ff9ab6d9e7bcc043594cacdbd432-VerificationService").(p9196ff9ab.VerificationService)
}


func (inst*p623d365faa_authx_AuthWithSMS) getAuthorSer(ie application.InjectionExt)p9196ff9ab.AuthorizationService{
    return ie.GetComponent("#alias-9196ff9ab6d9e7bcc043594cacdbd432-AuthorizationService").(p9196ff9ab.AuthorizationService)
}


func (inst*p623d365faa_authx_AuthWithSMS) getDBAgent(ie application.InjectionExt)p512a30914.Agent{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent").(p512a30914.Agent)
}


func (inst*p623d365faa_authx_AuthWithSMS) getUsers(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}


func (inst*p623d365faa_authx_AuthWithSMS) getEmailAddresses(ie application.InjectionExt)p2dece1e49.EmailAddressDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-EmailAddressDAO").(p2dece1e49.EmailAddressDAO)
}


func (inst*p623d365faa_authx_AuthWithSMS) getPhoneNumbers(ie application.InjectionExt)p2dece1e49.PhoneNumberDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PhoneNumberDAO").(p2dece1e49.PhoneNumberDAO)
}



// type p623d365fa.DefaultAuthorizer in package:github.com/starter-go/module-security-gin-gorm/internal/core/authx
//
// id:com-623d365faaa9ccfd-authx-DefaultAuthorizer
// class:
// alias:alias-9196ff9ab6d9e7bcc043594cacdbd432-AuthorizationService
// scope:singleton
//
type p623d365faa_authx_DefaultAuthorizer struct {
}

func (inst* p623d365faa_authx_DefaultAuthorizer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-623d365faaa9ccfd-authx-DefaultAuthorizer"
	r.Classes = ""
	r.Aliases = "alias-9196ff9ab6d9e7bcc043594cacdbd432-AuthorizationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p623d365faa_authx_DefaultAuthorizer) new() any {
    return &p623d365fa.DefaultAuthorizer{}
}

func (inst* p623d365faa_authx_DefaultAuthorizer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p623d365fa.DefaultAuthorizer)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.TokenMaxAge = inst.getTokenMaxAge(ie)
    com.SessionMaxAge = inst.getSessionMaxAge(ie)


    return nil
}


func (inst*p623d365faa_authx_DefaultAuthorizer) getService(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


func (inst*p623d365faa_authx_DefaultAuthorizer) getTokenMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${security.jwt.max-age}")
}


func (inst*p623d365faa_authx_DefaultAuthorizer) getSessionMaxAge(ie application.InjectionExt)int{
    return ie.GetInt("${security.session.max-age}")
}



// type p623d365fa.VerificationServiceImpl in package:github.com/starter-go/module-security-gin-gorm/internal/core/authx
//
// id:com-623d365faaa9ccfd-authx-VerificationServiceImpl
// class:
// alias:alias-9196ff9ab6d9e7bcc043594cacdbd432-VerificationService
// scope:singleton
//
type p623d365faa_authx_VerificationServiceImpl struct {
}

func (inst* p623d365faa_authx_VerificationServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-623d365faaa9ccfd-authx-VerificationServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-9196ff9ab6d9e7bcc043594cacdbd432-VerificationService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p623d365faa_authx_VerificationServiceImpl) new() any {
    return &p623d365fa.VerificationServiceImpl{}
}

func (inst* p623d365faa_authx_VerificationServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p623d365fa.VerificationServiceImpl)
	nop(ie, com)

	


    return nil
}



// type pcb7df5f5e.Convertor in package:github.com/starter-go/module-security-gin-gorm/internal/core/permissions
//
// id:com-cb7df5f5e18a8c87-permissions-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor
// scope:singleton
//
type pcb7df5f5e1_permissions_Convertor struct {
}

func (inst* pcb7df5f5e1_permissions_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cb7df5f5e18a8c87-permissions-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcb7df5f5e1_permissions_Convertor) new() any {
    return &pcb7df5f5e.Convertor{}
}

func (inst* pcb7df5f5e1_permissions_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcb7df5f5e.Convertor)
	nop(ie, com)

	


    return nil
}



// type pcb7df5f5e.Service in package:github.com/starter-go/module-security-gin-gorm/internal/core/permissions
//
// id:com-cb7df5f5e18a8c87-permissions-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-PermissionService
// scope:singleton
//
type pcb7df5f5e1_permissions_Service struct {
}

func (inst* pcb7df5f5e1_permissions_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cb7df5f5e18a8c87-permissions-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-PermissionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcb7df5f5e1_permissions_Service) new() any {
    return &pcb7df5f5e.Service{}
}

func (inst* pcb7df5f5e1_permissions_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcb7df5f5e.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*pcb7df5f5e1_permissions_Service) getDao(ie application.InjectionExt)p2dece1e49.PermissionDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionDAO").(p2dece1e49.PermissionDAO)
}


func (inst*pcb7df5f5e1_permissions_Service) getConvertor(ie application.InjectionExt)p2dece1e49.PermissionConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-PermissionConvertor").(p2dece1e49.PermissionConvertor)
}



// type pd8c903054.Convertor in package:github.com/starter-go/module-security-gin-gorm/internal/core/roles
//
// id:com-d8c903054a0aeb26-roles-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor
// scope:singleton
//
type pd8c903054a_roles_Convertor struct {
}

func (inst* pd8c903054a_roles_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d8c903054a0aeb26-roles-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd8c903054a_roles_Convertor) new() any {
    return &pd8c903054.Convertor{}
}

func (inst* pd8c903054a_roles_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd8c903054.Convertor)
	nop(ie, com)

	


    return nil
}



// type pd8c903054.Service in package:github.com/starter-go/module-security-gin-gorm/internal/core/roles
//
// id:com-d8c903054a0aeb26-roles-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-RoleService
// scope:singleton
//
type pd8c903054a_roles_Service struct {
}

func (inst* pd8c903054a_roles_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d8c903054a0aeb26-roles-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-RoleService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd8c903054a_roles_Service) new() any {
    return &pd8c903054.Service{}
}

func (inst* pd8c903054a_roles_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd8c903054.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*pd8c903054a_roles_Service) getDao(ie application.InjectionExt)p2dece1e49.RoleDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleDAO").(p2dece1e49.RoleDAO)
}


func (inst*pd8c903054a_roles_Service) getConvertor(ie application.InjectionExt)p2dece1e49.RoleConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-RoleConvertor").(p2dece1e49.RoleConvertor)
}



// type p12d48b32e.Service in package:github.com/starter-go/module-security-gin-gorm/internal/core/sessions
//
// id:com-12d48b32e189c661-sessions-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-SessionService
// scope:singleton
//
type p12d48b32e1_sessions_Service struct {
}

func (inst* p12d48b32e1_sessions_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-12d48b32e189c661-sessions-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-SessionService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p12d48b32e1_sessions_Service) new() any {
    return &p12d48b32e.Service{}
}

func (inst* p12d48b32e1_sessions_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p12d48b32e.Service)
	nop(ie, com)

	
    com.JWTser = inst.getJWTser(ie)


    return nil
}


func (inst*p12d48b32e1_sessions_Service) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}



// type p6bb98f719.JWTCodec in package:github.com/starter-go/module-security-gin-gorm/internal/core/tokens
//
// id:com-6bb98f71931b600f-tokens-JWTCodec
// class:class-91f218d46ec21cd234778bbe54aecc66-Registry
// alias:
// scope:singleton
//
type p6bb98f7193_tokens_JWTCodec struct {
}

func (inst* p6bb98f7193_tokens_JWTCodec) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6bb98f71931b600f-tokens-JWTCodec"
	r.Classes = "class-91f218d46ec21cd234778bbe54aecc66-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6bb98f7193_tokens_JWTCodec) new() any {
    return &p6bb98f719.JWTCodec{}
}

func (inst* p6bb98f7193_tokens_JWTCodec) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6bb98f719.JWTCodec)
	nop(ie, com)

	
    com.Secret = inst.getSecret(ie)


    return nil
}


func (inst*p6bb98f7193_tokens_JWTCodec) getSecret(ie application.InjectionExt)string{
    return ie.GetString("${security.jwt.secret}")
}



// type p217db24b0.Convertor in package:github.com/starter-go/module-security-gin-gorm/internal/core/users
//
// id:com-217db24b0092058d-users-Convertor
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor
// scope:singleton
//
type p217db24b00_users_Convertor struct {
}

func (inst* p217db24b00_users_Convertor) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-217db24b0092058d-users-Convertor"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p217db24b00_users_Convertor) new() any {
    return &p217db24b0.Convertor{}
}

func (inst* p217db24b00_users_Convertor) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p217db24b0.Convertor)
	nop(ie, com)

	


    return nil
}



// type p217db24b0.Service in package:github.com/starter-go/module-security-gin-gorm/internal/core/users
//
// id:com-217db24b0092058d-users-Service
// class:
// alias:alias-2dece1e495fd61b93f78009d229f38cf-UserService
// scope:singleton
//
type p217db24b00_users_Service struct {
}

func (inst* p217db24b00_users_Service) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-217db24b0092058d-users-Service"
	r.Classes = ""
	r.Aliases = "alias-2dece1e495fd61b93f78009d229f38cf-UserService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p217db24b00_users_Service) new() any {
    return &p217db24b0.Service{}
}

func (inst* p217db24b00_users_Service) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p217db24b0.Service)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Convertor = inst.getConvertor(ie)


    return nil
}


func (inst*p217db24b00_users_Service) getDao(ie application.InjectionExt)p2dece1e49.UserDAO{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserDAO").(p2dece1e49.UserDAO)
}


func (inst*p217db24b00_users_Service) getConvertor(ie application.InjectionExt)p2dece1e49.UserConvertor{
    return ie.GetComponent("#alias-2dece1e495fd61b93f78009d229f38cf-UserConvertor").(p2dece1e49.UserConvertor)
}


