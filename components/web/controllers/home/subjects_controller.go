package home

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
)

// SubjectVO ...
type SubjectVO struct {
	rbac.SubjectVO
}

////////////////////////////////////////////////////////////////////////////////

// SubjectController ...
type SubjectController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder    //starter:inject("#")
	Service   rbac.SubjectService //starter:inject("#")

}

func (inst *SubjectController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *SubjectController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *SubjectController) route(g libgin.RouterProxy) error {
	g = g.For("subject")
	g.GET("", inst.handleGetCurrent)
	return nil
}

func (inst *SubjectController) execute(req *mySubjectRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *SubjectController) handleGetCurrent(c *gin.Context) {
	req := &mySubjectRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetCurrentSubject)
}

////////////////////////////////////////////////////////////////////////////////

type mySubjectRequest struct {
	// contexts
	controller *SubjectController
	context    *gin.Context

	// flags
	wantRequestBody bool
	wantRequestID   bool
	wantRequestRBAC bool

	// params
	id    rbac.SubjectID
	roles rbac.RoleNameList

	// body
	body1 SubjectVO
	body2 SubjectVO
}

func (inst *mySubjectRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestID {
		str := c.Param("id")
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		inst.id = rbac.SubjectID(n)
	}

	return nil
}

func (inst *mySubjectRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *mySubjectRequest) doGetCurrentSubject() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := ser.GetCurrent(ctx)
	if err != nil {
		return err
	}
	if o1 == nil {
		o1 = &rbac.SubjectDTO{}
	}
	inst.body2.Subject = o1
	return nil
}

////////////////////////////////////////////////////////////////////////////////
