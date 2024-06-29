package main4ggsecurity

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

    
    inst.register(&p04a90047b8_signup_AuthorizerForSignUp{})
    inst.register(&p16ea5b788f_home_AuthController{})
    inst.register(&p16ea5b788f_home_ExampleController{})
    inst.register(&p16ea5b788f_home_PermissionController{})
    inst.register(&p16ea5b788f_home_RegionController{})
    inst.register(&p16ea5b788f_home_RoleController{})
    inst.register(&p16ea5b788f_home_SessionController{})
    inst.register(&p16ea5b788f_home_SubjectController{})
    inst.register(&p16ea5b788f_home_TokenController{})
    inst.register(&p16ea5b788f_home_UserController{})
    inst.register(&p492cd0c9ee_sms_AuthBySMS{})
    inst.register(&p4f3ce922c5_email_AuthByEmail{})
    inst.register(&p56e4dad453_services_PermissionImportServiceImpl{})
    inst.register(&p775fb156f4_develop_DebugController{})
    inst.register(&p775fb156f4_develop_ExampleController{})
    inst.register(&p89ee959c3d_login_AuthorizerForLogin{})
    inst.register(&p985949a7f7_admin_GroupController{})
    inst.register(&p985949a7f7_admin_PermissionController{})
    inst.register(&p985949a7f7_admin_RegionController{})
    inst.register(&p985949a7f7_admin_RoleController{})
    inst.register(&p985949a7f7_admin_UserController{})
    inst.register(&pc0223a2c4a_vericode_VerificationServiceImpl{})


    return nil
}
