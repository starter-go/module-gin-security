package gen4ggsecurity

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p13ed4c221a_develop_DebugController{})
    inst.register(&p13ed4c221a_develop_ExampleController{})
    inst.register(&p736bdaa375_home_AuthController{})
    inst.register(&p736bdaa375_home_ExampleController{})
    inst.register(&p736bdaa375_home_PermissionController{})
    inst.register(&p736bdaa375_home_RegionController{})
    inst.register(&p736bdaa375_home_RoleController{})
    inst.register(&p736bdaa375_home_SessionController{})
    inst.register(&p736bdaa375_home_TokenController{})
    inst.register(&p736bdaa375_home_UserController{})
    inst.register(&p7f9a7d1cfc_signup_AuthorizerForSignUp{})
    inst.register(&pbd5217dfae_vericode_VerificationServiceImpl{})
    inst.register(&pd4be75b3b4_sms_AuthBySMS{})
    inst.register(&pdc51ff5e4d_login_AuthorizerForLogin{})
    inst.register(&pef5fb8e156_admin_PermissionController{})
    inst.register(&pef5fb8e156_admin_RegionController{})
    inst.register(&pef5fb8e156_admin_RoleController{})
    inst.register(&pef5fb8e156_admin_UserController{})
    inst.register(&pf48f932fb4_email_AuthByEmail{})


    return nil
}
