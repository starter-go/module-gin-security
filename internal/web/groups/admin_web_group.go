package groups

import (
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/implements"
)

// AdminWebGroup ...
type AdminWebGroup struct {

	//starter:component
	_as func(libgin.Group) //starter:as(".")

	Context libgin.Context //starter:inject("#")
	Name    string         //starter:inject("${web-group.admin.name}")
	Path    string         //starter:inject("${web-group.admin.path}")

	inner libgin.Group
}

func (inst *AdminWebGroup) _impl() libgin.Group {
	return inst
}

// Registration ...
func (inst *AdminWebGroup) Registration() *libgin.GroupRegistration {

	i := inst.inner
	if i == nil {
		i = inst.loadInner()
		inst.inner = i
	}

	return &libgin.GroupRegistration{
		Name:  inst.Name,
		Path:  inst.Path,
		Group: inst,
	}
}

// Route ...
func (inst *AdminWebGroup) Route(r libgin.Router) error {
	return inst.inner.Route(r)
}

func (inst *AdminWebGroup) loadInner() libgin.Group {
	gb := &implements.GroupBuilder{
		Context: inst.Context,
		Name:    inst.Name,
		Path:    inst.Path,
	}
	return gb.Create()
}
