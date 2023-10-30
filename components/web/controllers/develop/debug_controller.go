package develop

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/rbac"
)

// DebugVO ...
type DebugVO struct {
	rbac.BaseVO

	Items []*rbac.PermissionDTO `json:"items"`
}

////////////////////////////////////////////////////////////////////////////////

// DebugController ...
type DebugController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder       //starter:inject("#")
	Service   rbac.PermissionService //starter:inject("#")

}

func (inst *DebugController) _impl(a application.Lifecycle) {
	inst._as(inst)
	a = inst
}

// Life ...
func (inst *DebugController) Life() *application.Life {
	return &application.Life{
		OnStart: inst.onStart,
	}
}

func (inst *DebugController) onStart() error {

	cache := inst.Service.GetCache()
	cache.Clear()

	return nil
}

// Registration ...
func (inst *DebugController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Groups: []string{"develop"},
		Route:  inst.route,
	}
}

func (inst *DebugController) route(g libgin.RouterProxy) error {
	g = g.For("debug")
	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)
	return nil
}

func (inst *DebugController) execute(req *myDebugRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *DebugController) handleGetList(c *gin.Context) {
	req := &myDebugRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *DebugController) handleGetOne(c *gin.Context) {
	req := &myDebugRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *DebugController) handleInsert(c *gin.Context) {
	req := &myDebugRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *DebugController) handleUpdate(c *gin.Context) {
	req := &myDebugRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *DebugController) handleRemove(c *gin.Context) {
	req := &myDebugRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myDebugRequest struct {
	// contexts
	controller *DebugController
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
	body1 DebugVO
	body2 DebugVO
}

func (inst *myDebugRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myDebugRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myDebugRequest) doGetList() error {
	return nil
}

func (inst *myDebugRequest) doGetOne() error {
	return nil
}

func (inst *myDebugRequest) doInsert() error {
	return nil
}

func (inst *myDebugRequest) doUpdate() error {
	return nil
}

func (inst *myDebugRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
