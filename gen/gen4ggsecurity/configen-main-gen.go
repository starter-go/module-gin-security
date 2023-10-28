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

    
    inst.register(&p736bdaa375_home_AuthController{})
    inst.register(&p736bdaa375_home_ExampleController{})
    inst.register(&p736bdaa375_home_PermissionController{})
    inst.register(&p736bdaa375_home_RoleController{})
    inst.register(&p736bdaa375_home_SessionController{})
    inst.register(&p736bdaa375_home_TokenController{})
    inst.register(&p736bdaa375_home_UserController{})
    inst.register(&pef5fb8e156_admin_PermissionController{})
    inst.register(&pef5fb8e156_admin_RoleController{})
    inst.register(&pef5fb8e156_admin_UserController{})


    return nil
}
