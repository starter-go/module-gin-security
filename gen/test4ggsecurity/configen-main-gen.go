package test4ggsecurity

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

    
    inst.register(&pe7f155628f_testcases_APITester{})
    inst.register(&pe7f155628f_testcases_AuthSendVerificationMailTester{})
    inst.register(&pe7f155628f_testcases_AuthSendVerificationSMSTester{})
    inst.register(&pe7f155628f_testcases_AuthSignInTester{})
    inst.register(&pe7f155628f_testcases_AuthSignUpTester{})
    inst.register(&pe7f155628f_testcases_DefaultResponseHandler{})
    inst.register(&pe7f155628f_testcases_MockEmailSender{})
    inst.register(&pe7f155628f_testcases_MockSMSSender{})
    inst.register(&pe7f155628f_testcases_MyHTTPAgentFilter{})
    inst.register(&pe7f155628f_testcases_SessionsCurrentTester{})


    return nil
}
