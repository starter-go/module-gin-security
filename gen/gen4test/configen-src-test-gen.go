package gen4test
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p2c7ffb218 "github.com/starter-go/module-gin-security/src/test/golang"
    p91f218d46 "github.com/starter-go/security/jwt"
     "github.com/starter-go/application"
)

// type p2c7ffb218.Demo1Controller in package:github.com/starter-go/module-gin-security/src/test/golang
//
// id:com-2c7ffb21840730eb-golang-Demo1Controller
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p2c7ffb2184_golang_Demo1Controller struct {
}

func (inst* p2c7ffb2184_golang_Demo1Controller) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2c7ffb21840730eb-golang-Demo1Controller"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2c7ffb2184_golang_Demo1Controller) new() any {
    return &p2c7ffb218.Demo1Controller{}
}

func (inst* p2c7ffb2184_golang_Demo1Controller) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2c7ffb218.Demo1Controller)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)
    com.JWTser = inst.getJWTser(ie)


    return nil
}


func (inst*p2c7ffb2184_golang_Demo1Controller) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p2c7ffb2184_golang_Demo1Controller) getJWTser(ie application.InjectionExt)p91f218d46.Service{
    return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}


