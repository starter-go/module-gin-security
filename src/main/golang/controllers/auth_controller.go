package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/rbac"
)

// AuthVO ...
type AuthVO struct {
	rbac.BaseVO

	Auth rbac.AuthDTO `json:"auth"`
}

////////////////////////////////////////////////////////////////////////////////

// AuthController ...
type AuthController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service   rbac.AuthService //starter:inject("#")

}

func (inst *AuthController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *AuthController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *AuthController) route(g *gin.RouterGroup) error {
	g = g.Group("auth")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handlePost)

	return nil
}

func (inst *AuthController) execute(req *myAuthRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *AuthController) handleGetList(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *AuthController) handleGetOne(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *AuthController) handlePost(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doLogin)
}

func (inst *AuthController) handleUpdate(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *AuthController) handleRemove(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myAuthRequest struct {
	// contexts
	controller *AuthController
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
	body1 AuthVO
	body2 AuthVO
}

func (inst *myAuthRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myAuthRequest) doGetList() error {
	return nil
}

func (inst *myAuthRequest) doGetOne() error {
	return nil
}

func (inst *myAuthRequest) doLogin() error {

	ctx := inst.context
	ser := inst.controller.Service
	o1 := &inst.body1.Auth

	o2, err := ser.Login(ctx, o1)
	if err != nil {
		return err
	}

	inst.body2.Auth = *o2
	return nil
}

func (inst *myAuthRequest) doUpdate() error {
	return nil
}

func (inst *myAuthRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
