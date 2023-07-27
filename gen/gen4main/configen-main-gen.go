package gen4main

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

    
    inst.register(&p4ba0ff4dda_tokens_GinJWTokenAdapter{})
    inst.register(&pda2eabb8fd_controllers_AuthController{})
    inst.register(&pda2eabb8fd_controllers_ExampleController{})
    inst.register(&pda2eabb8fd_controllers_PermissionController{})
    inst.register(&pda2eabb8fd_controllers_RoleController{})
    inst.register(&pda2eabb8fd_controllers_SessionController{})
    inst.register(&pda2eabb8fd_controllers_UserController{})
    inst.register(&ped246966fe_golang_JSONGinResponder{})


    return nil
}
