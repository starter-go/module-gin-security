package home

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
	"github.com/starter-go/security/subjects"
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

	Responder libgin.Responder       //starter:inject("#")
	Service   rbac.PermissionService //starter:inject("#")

	// JWTService jwt.Service            //starter:inject("#")

}

func (inst *TokenController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *TokenController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *TokenController) route(g libgin.RouterProxy) error {
	g = g.For("tokens")

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
	sub, err := subjects.Current(ctx)
	if err != nil {
		return err
	}
	token := sub.GetToken()
	o1 := token.Get()
	inst.body2.Tokens = []*rbac.TokenDTO{o1}
	return nil
}

func (inst *myTokenRequest) doInsert() error {
	return nil
}

func (inst *myTokenRequest) doRenew() error {

	ctx := inst.context
	sub, err := subjects.Current(ctx)
	if err != nil {
		return err
	}

	// params
	list1 := inst.body1.Tokens
	item1, err := controllers.GetItemOnlyOne[rbac.TokenDTO](list1)
	if err != nil {
		return err
	}

	t1 := lang.Now()
	t2 := item1.ExpiredAt
	if t2 == 0 {
		t2 = t1.Add(120 * time.Second) // 默认 max-age 为 120 s
	}

	// older
	token := sub.GetToken()
	tk1 := token.Get()

	// newer
	tk2 := &rbac.TokenDTO{}
	*tk2 = *tk1
	tk2.StartedAt = t1
	tk2.ExpiredAt = t2

	// update
	token.Set(tk2)
	err = token.Commit()
	if err != nil {
		return err
	}

	// done
	inst.body2.Tokens = []*rbac.TokenDTO{tk2}
	return nil
}

func (inst *myTokenRequest) doUpdate() error {
	return nil
}

func (inst *myTokenRequest) doRemove() error {
	return nil
}

// func (inst *myTokenRequest) castTokenInfo(src *jwt.Token, dst *rbac.TokenDTO) {
// 	dst.CreatedAt = src.CreatedAt
// 	dst.ExpiredAt = src.ExpiredAt
// 	dst.MaxAge = src.MaxAge
// 	dst.Session = &src.Session
// }

////////////////////////////////////////////////////////////////////////////////
