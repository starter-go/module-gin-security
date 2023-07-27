package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/rbac"
)

// SessionVO ...
type SessionVO struct {
	rbac.BaseVO

	Sessions []*rbac.SessionDTO `json:"sessions"`
}

////////////////////////////////////////////////////////////////////////////////

// SessionController ...
type SessionController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder       //starter:inject("#")
	Service   rbac.PermissionService //starter:inject("#")

}

func (inst *SessionController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *SessionController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *SessionController) route(g *gin.RouterGroup) error {
	g = g.Group("sessions")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *SessionController) execute(req *mySessionRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *SessionController) handleGetList(c *gin.Context) {
	req := &mySessionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *SessionController) handleGetOne(c *gin.Context) {
	req := &mySessionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *SessionController) handleInsert(c *gin.Context) {
	req := &mySessionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *SessionController) handleUpdate(c *gin.Context) {
	req := &mySessionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *SessionController) handleRemove(c *gin.Context) {
	req := &mySessionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type mySessionRequest struct {
	// contexts
	controller *SessionController
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
	body1 SessionVO
	body2 SessionVO
}

func (inst *mySessionRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySessionRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *mySessionRequest) doGetList() error {
	return nil
}

func (inst *mySessionRequest) doGetOne() error {
	return nil
}

func (inst *mySessionRequest) doInsert() error {
	return nil
}

func (inst *mySessionRequest) doUpdate() error {
	return nil
}

func (inst *mySessionRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
