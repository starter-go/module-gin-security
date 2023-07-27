package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/rbac"
)

// RoleVO ...
type RoleVO struct {
	rbac.BaseVO

	Roles []*rbac.RoleDTO `json:"roles"`
}

////////////////////////////////////////////////////////////////////////////////

// RoleController ...
type RoleController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service   rbac.RoleService //starter:inject("#")

}

func (inst *RoleController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *RoleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *RoleController) route(g *gin.RouterGroup) error {
	g = g.Group("roles")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *RoleController) execute(req *myRoleRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *RoleController) handleGetList(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *RoleController) handleGetOne(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *RoleController) handleInsert(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *RoleController) handleUpdate(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *RoleController) handleRemove(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myRoleRequest struct {
	// contexts
	controller *RoleController
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
	body1 RoleVO
	body2 RoleVO
}

func (inst *myRoleRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRoleRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myRoleRequest) doGetList() error {
	return nil
}

func (inst *myRoleRequest) doGetOne() error {
	return nil
}

func (inst *myRoleRequest) doInsert() error {
	return nil
}

func (inst *myRoleRequest) doUpdate() error {
	return nil
}

func (inst *myRoleRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
