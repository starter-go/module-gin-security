package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/rbac"
)

// PermissionVO ...
type PermissionVO struct {
	rbac.BaseVO

	Permissions []*rbac.PermissionDTO `json:"permissions"`
}

////////////////////////////////////////////////////////////////////////////////

// PermissionController ...
type PermissionController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder       //starter:inject("#")
	Service   rbac.PermissionService //starter:inject("#")

}

func (inst *PermissionController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *PermissionController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *PermissionController) route(g *gin.RouterGroup) error {
	g = g.Group("permissions")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *PermissionController) execute(req *myPermissionRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *PermissionController) handleGetList(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *PermissionController) handleGetOne(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *PermissionController) handleInsert(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *PermissionController) handleUpdate(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *PermissionController) handleRemove(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myPermissionRequest struct {
	// contexts
	controller *PermissionController
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
	body1 PermissionVO
	body2 PermissionVO
}

func (inst *myPermissionRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPermissionRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myPermissionRequest) doGetList() error {
	return nil
}

func (inst *myPermissionRequest) doGetOne() error {
	return nil
}

func (inst *myPermissionRequest) doInsert() error {
	return nil
}

func (inst *myPermissionRequest) doUpdate() error {
	return nil
}

func (inst *myPermissionRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
