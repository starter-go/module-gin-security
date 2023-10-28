package impl

import "github.com/starter-go/libgin"

// ExampleController 。。。
type ExampleController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")
}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ExampleController) route(g libgin.RouterProxy) error {
	g = g.For("example")
	return nil
}
