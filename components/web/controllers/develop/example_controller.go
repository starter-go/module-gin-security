package develop

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
)

// ExampleVO ...
type ExampleVO struct {
	rbac.BaseVO

	Items []*rbac.PermissionDTO `json:"items"`
}

////////////////////////////////////////////////////////////////////////////////

// ExampleController ...
type ExampleController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder       //starter:inject("#")
	Service   rbac.PermissionService //starter:inject("#")

}

func (inst *ExampleController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{"develop"},
		Route:  inst.route,
	}
}

func (inst *ExampleController) route(g libgin.RouterProxy) error {
	g = g.For("Example")
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)
	return nil
}

func (inst *ExampleController) execute(req *myExampleRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *ExampleController) handleGetList(c *gin.Context) {
	req := &myExampleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *ExampleController) handleGetOne(c *gin.Context) {
	req := &myExampleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *ExampleController) handleInsert(c *gin.Context) {
	req := &myExampleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *ExampleController) handleUpdate(c *gin.Context) {
	req := &myExampleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *ExampleController) handleRemove(c *gin.Context) {
	req := &myExampleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleRequest struct {
	// contexts
	controller *ExampleController
	context    *gin.Context

	// flags
	wantRequestBody bool
	wantRequestID   bool
	wantRequestPage bool
	wantRequestRBAC bool

	// params
	pagination rbac.Pagination
	id         rbac.PermissionID
	roles      rbac.RoleNameList

	// body
	body1 ExampleVO
	body2 ExampleVO
}

func (inst *myExampleRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myExampleRequest) doGetList() error {
	return nil
}

func (inst *myExampleRequest) doGetOne() error {
	return nil
}

func (inst *myExampleRequest) doInsert() error {
	return nil
}

func (inst *myExampleRequest) doUpdate() error {
	return nil
}

func (inst *myExampleRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
