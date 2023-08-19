package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/rbac"
)

// TokenVO ...
type TokenVO struct {
	rbac.BaseVO

	Tokens []*rbac.TokenDTO `json:"tokens"`
}

////////////////////////////////////////////////////////////////////////////////

// TokenController ...
type TokenController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder  libgin.Responder       //starter:inject("#")
	Service    rbac.PermissionService //starter:inject("#")
	JWTService jwt.Service            //starter:inject("#")

}

func (inst *TokenController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *TokenController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *TokenController) route(g *gin.RouterGroup) error {
	g = g.Group("tokens")

	g.GET("", inst.handleGetList)
	g.GET("current", inst.handleGetCurrent)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)
	g.POST("renew", inst.handleRenew)

	return nil
}

func (inst *TokenController) execute(req *myTokenRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *TokenController) handleGetList(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *TokenController) handleGetOne(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *TokenController) handleGetCurrent(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetCurrent)
}

func (inst *TokenController) handleInsert(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *TokenController) handleRenew(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRenew)
}

func (inst *TokenController) handleUpdate(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *TokenController) handleRemove(c *gin.Context) {
	req := &myTokenRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myTokenRequest struct {
	// contexts
	controller *TokenController
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
	body1 TokenVO
	body2 TokenVO
}

func (inst *myTokenRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myTokenRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myTokenRequest) doGetList() error {
	return nil
}

func (inst *myTokenRequest) doGetOne() error {
	return nil
}

func (inst *myTokenRequest) doGetCurrent() error {
	ctx := inst.context
	t1, err := inst.controller.JWTService.GetDTO(ctx)
	t2 := &rbac.TokenDTO{}
	if err == nil {
		inst.castTokenInfo(t1, t2)
	}
	list := inst.body2.Tokens
	list = append(list, t2)
	inst.body2.Tokens = list
	return nil
}

func (inst *myTokenRequest) doInsert() error {
	return nil
}

func (inst *myTokenRequest) doRenew() error {

	ctx := inst.context
	ser := inst.controller.JWTService

	// older
	t1, err := ser.GetDTO(ctx)
	if err != nil {
		return err
	}

	// newer
	t2 := &jwt.Token{}
	t2.Session = t1.Session
	err = ser.SetDTO(ctx, t2)
	if err != nil {
		return err
	}

	// response
	t3 := &rbac.TokenDTO{}
	inst.castTokenInfo(t2, t3)
	list := inst.body2.Tokens
	list = append(list, t3)
	inst.body2.Tokens = list
	return nil
}

func (inst *myTokenRequest) doUpdate() error {
	return nil
}

func (inst *myTokenRequest) doRemove() error {
	return nil
}

func (inst *myTokenRequest) castTokenInfo(src *jwt.Token, dst *rbac.TokenDTO) {
	dst.CreatedAt = src.CreatedAt
	dst.ExpiredAt = src.ExpiredAt
	dst.Session = &src.Session
}

////////////////////////////////////////////////////////////////////////////////
