package home

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
	"github.com/starter-go/security/rbac"
)

// UserVO ...
type UserVO struct {
	rbac.UserVO

	// rbac.BaseVO
	// Users []*rbac.UserDTO `json:"users"`
}

////////////////////////////////////////////////////////////////////////////////

// UserController ...
type UserController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service   rbac.UserService //starter:inject("#")

}

func (inst *UserController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *UserController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *UserController) route(g libgin.RouterProxy) error {
	g = g.For("users")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *UserController) execute(req *myUserRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *UserController) handleGetList(c *gin.Context) {
	req := &myUserRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: true,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *UserController) handleGetOne(c *gin.Context) {
	req := &myUserRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *UserController) handleInsert(c *gin.Context) {
	req := &myUserRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *UserController) handleUpdate(c *gin.Context) {
	req := &myUserRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *UserController) handleRemove(c *gin.Context) {
	req := &myUserRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myUserRequest struct {
	// contexts
	controller *UserController
	context    *gin.Context

	// flags
	wantRequestBody bool
	wantRequestID   bool
	wantRequestPage bool
	wantRequestRBAC bool

	// params
	id         rbac.UserID
	pagination rbac.Pagination
	roles      rbac.RoleNameList

	// body
	body1 UserVO
	body2 UserVO
}

func (inst *myUserRequest) open() error {

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
		inst.id = rbac.UserID(n)
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
	}

	return nil
}

func (inst *myUserRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myUserRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	page := &inst.pagination
	q := &rbac.UserQuery{}
	q.Pagination = *page
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Users = list
	inst.body2.Pagination = page
	return nil
}

func (inst *myUserRequest) doGetOne() error {
	return nil
}

func (inst *myUserRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := controllers.GetItemOnlyOne[rbac.UserDTO](inst.body1.Users)
	if err != nil {
		return err
	}
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Users = []*rbac.UserDTO{o2}
	return nil
}

func (inst *myUserRequest) doUpdate() error {
	return nil
}

func (inst *myUserRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
