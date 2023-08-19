package gen4data

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

    
    inst.register(&p4879aa5e6a_data_EmailDao{})
    inst.register(&p4879aa5e6a_data_PermissionDao{})
    inst.register(&p4879aa5e6a_data_PhoneDao{})
    inst.register(&p4879aa5e6a_data_RoleDao{})
    inst.register(&p4879aa5e6a_data_UserDao{})


    return nil
}
