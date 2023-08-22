package gen4web

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

    
    inst.register(&pa70cfe71fb_admin_PermissionController{})
    inst.register(&pa70cfe71fb_admin_RoleController{})
    inst.register(&pa70cfe71fb_admin_UserController{})
    inst.register(&pa9413a1971_web_GinJWTokenAdapter{})
    inst.register(&pa9413a1971_web_JSONGinResponder{})
    inst.register(&pe1f113ef5f_controllers_AuthController{})
    inst.register(&pe1f113ef5f_controllers_ExampleController{})
    inst.register(&pe1f113ef5f_controllers_PermissionController{})
    inst.register(&pe1f113ef5f_controllers_RoleController{})
    inst.register(&pe1f113ef5f_controllers_SessionController{})
    inst.register(&pe1f113ef5f_controllers_TokenController{})
    inst.register(&pe1f113ef5f_controllers_UserController{})
    inst.register(&pfbdcfa5284_groups_AdminWebGroup{})


    return nil
}
