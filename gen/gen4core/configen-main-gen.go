package gen4core

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

    
    inst.register(&p12d48b32e1_sessions_Service{})
    inst.register(&p217db24b00_users_Convertor{})
    inst.register(&p217db24b00_users_Service{})
    inst.register(&p623d365faa_authx_AuthWithEmail{})
    inst.register(&p623d365faa_authx_AuthWithSMS{})
    inst.register(&p623d365faa_authx_DefaultAuthorizer{})
    inst.register(&p623d365faa_authx_PasswordAuth{})
    inst.register(&p623d365faa_authx_VerificationServiceImpl{})
    inst.register(&p6bb98f7193_tokens_JWTCodec{})
    inst.register(&pcb7df5f5e1_permissions_Convertor{})
    inst.register(&pcb7df5f5e1_permissions_Service{})
    inst.register(&pd8c903054a_roles_Convertor{})
    inst.register(&pd8c903054a_roles_Service{})


    return nil
}
