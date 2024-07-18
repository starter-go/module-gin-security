package test4ggsecurity
import (
    pdea5a0f47 "github.com/starter-go/httpagent"
    pe7f155628 "github.com/starter-go/security-gin-gorm/src/test/golang/testcases"
     "github.com/starter-go/application"
)

// type pe7f155628.APITester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-APITester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_APITester struct {
}

func (inst* pe7f155628f_testcases_APITester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-APITester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_APITester) new() any {
    return &pe7f155628.APITester{}
}

func (inst* pe7f155628f_testcases_APITester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.APITester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_APITester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_APITester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.AuthSendVerificationMailTester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-AuthSendVerificationMailTester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_AuthSendVerificationMailTester struct {
}

func (inst* pe7f155628f_testcases_AuthSendVerificationMailTester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-AuthSendVerificationMailTester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_AuthSendVerificationMailTester) new() any {
    return &pe7f155628.AuthSendVerificationMailTester{}
}

func (inst* pe7f155628f_testcases_AuthSendVerificationMailTester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.AuthSendVerificationMailTester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_AuthSendVerificationMailTester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_AuthSendVerificationMailTester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.AuthSendVerificationSMSTester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-AuthSendVerificationSMSTester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_AuthSendVerificationSMSTester struct {
}

func (inst* pe7f155628f_testcases_AuthSendVerificationSMSTester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-AuthSendVerificationSMSTester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_AuthSendVerificationSMSTester) new() any {
    return &pe7f155628.AuthSendVerificationSMSTester{}
}

func (inst* pe7f155628f_testcases_AuthSendVerificationSMSTester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.AuthSendVerificationSMSTester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_AuthSendVerificationSMSTester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_AuthSendVerificationSMSTester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.AuthSignInTester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-AuthSignInTester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_AuthSignInTester struct {
}

func (inst* pe7f155628f_testcases_AuthSignInTester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-AuthSignInTester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_AuthSignInTester) new() any {
    return &pe7f155628.AuthSignInTester{}
}

func (inst* pe7f155628f_testcases_AuthSignInTester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.AuthSignInTester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_AuthSignInTester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_AuthSignInTester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.AuthSignUpTester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-AuthSignUpTester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_AuthSignUpTester struct {
}

func (inst* pe7f155628f_testcases_AuthSignUpTester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-AuthSignUpTester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_AuthSignUpTester) new() any {
    return &pe7f155628.AuthSignUpTester{}
}

func (inst* pe7f155628f_testcases_AuthSignUpTester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.AuthSignUpTester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_AuthSignUpTester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_AuthSignUpTester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.SessionsCurrentTester in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-SessionsCurrentTester
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe7f155628f_testcases_SessionsCurrentTester struct {
}

func (inst* pe7f155628f_testcases_SessionsCurrentTester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-SessionsCurrentTester"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_SessionsCurrentTester) new() any {
    return &pe7f155628.SessionsCurrentTester{}
}

func (inst* pe7f155628f_testcases_SessionsCurrentTester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.SessionsCurrentTester)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.ResponseHandler = inst.getResponseHandler(ie)


    return nil
}


func (inst*pe7f155628f_testcases_SessionsCurrentTester) getAgent(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}


func (inst*pe7f155628f_testcases_SessionsCurrentTester) getResponseHandler(ie application.InjectionExt)pe7f155628.ResponseHandler{
    return ie.GetComponent("#alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler").(pe7f155628.ResponseHandler)
}



// type pe7f155628.DefaultResponseHandler in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-DefaultResponseHandler
// class:
// alias:alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler
// scope:singleton
//
type pe7f155628f_testcases_DefaultResponseHandler struct {
}

func (inst* pe7f155628f_testcases_DefaultResponseHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-DefaultResponseHandler"
	r.Classes = ""
	r.Aliases = "alias-e7f155628f0ebd15880a930dee7b8781-ResponseHandler"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_DefaultResponseHandler) new() any {
    return &pe7f155628.DefaultResponseHandler{}
}

func (inst* pe7f155628f_testcases_DefaultResponseHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.DefaultResponseHandler)
	nop(ie, com)

	


    return nil
}



// type pe7f155628.MockEmailSender in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-MockEmailSender
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type pe7f155628f_testcases_MockEmailSender struct {
}

func (inst* pe7f155628f_testcases_MockEmailSender) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-MockEmailSender"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_MockEmailSender) new() any {
    return &pe7f155628.MockEmailSender{}
}

func (inst* pe7f155628f_testcases_MockEmailSender) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.MockEmailSender)
	nop(ie, com)

	


    return nil
}



// type pe7f155628.MockSMSSender in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-MockSMSSender
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type pe7f155628f_testcases_MockSMSSender struct {
}

func (inst* pe7f155628f_testcases_MockSMSSender) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-MockSMSSender"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_MockSMSSender) new() any {
    return &pe7f155628.MockSMSSender{}
}

func (inst* pe7f155628f_testcases_MockSMSSender) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.MockSMSSender)
	nop(ie, com)

	


    return nil
}



// type pe7f155628.MyHTTPAgentFilter in package:github.com/starter-go/security-gin-gorm/src/test/golang/testcases
//
// id:com-e7f155628f0ebd15-testcases-MyHTTPAgentFilter
// class:class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry
// alias:
// scope:singleton
//
type pe7f155628f_testcases_MyHTTPAgentFilter struct {
}

func (inst* pe7f155628f_testcases_MyHTTPAgentFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e7f155628f0ebd15-testcases-MyHTTPAgentFilter"
	r.Classes = "class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe7f155628f_testcases_MyHTTPAgentFilter) new() any {
    return &pe7f155628.MyHTTPAgentFilter{}
}

func (inst* pe7f155628f_testcases_MyHTTPAgentFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe7f155628.MyHTTPAgentFilter)
	nop(ie, com)

	


    return nil
}


