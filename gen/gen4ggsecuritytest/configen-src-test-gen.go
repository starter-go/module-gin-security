package gen4ggsecuritytest

import (
	"github.com/starter-go/application"
	pd1a916a20 "github.com/starter-go/libgin"
	pae18fcf34 "github.com/starter-go/security-gin-gorm/src/test/golang"
	p91f218d46 "github.com/starter-go/security/jwt"
)

// type pae18fcf34.Demo1Controller in package:github.com/starter-go/security-gin-gorm/src/test/golang
//
// id:com-ae18fcf347d8eb13-golang-Demo1Controller
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pae18fcf347_golang_Demo1Controller struct {
}

func (inst *pae18fcf347_golang_Demo1Controller) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ae18fcf347d8eb13-golang-Demo1Controller"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst *pae18fcf347_golang_Demo1Controller) new() any {
	return &pae18fcf34.Demo1Controller{}
}

func (inst *pae18fcf347_golang_Demo1Controller) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pae18fcf34.Demo1Controller)
	nop(ie, com)

	com.Responder = inst.getResponder(ie)
	com.JWTser = inst.getJWTser(ie)

	return nil
}

func (inst *pae18fcf347_golang_Demo1Controller) getResponder(ie application.InjectionExt) pd1a916a20.Responder {
	return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}

func (inst *pae18fcf347_golang_Demo1Controller) getJWTser(ie application.InjectionExt) p91f218d46.Service {
	return ie.GetComponent("#alias-91f218d46ec21cd234778bbe54aecc66-Service").(p91f218d46.Service)
}
