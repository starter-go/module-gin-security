package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
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
	return &libgin.ControllerRegistration{
		Route:  inst.route,
		Groups: []string{"admin"},
	}
}

func (inst *RoleController) route(g libgin.RouterProxy) error {
	g = g.For("roles")

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
		wantRequestID:   false,
		wantRequestPage: true,
	}
	inst.execute(req, req.doGetList)
}

func (inst *RoleController) handleGetOne(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *RoleController) handleInsert(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *RoleController) handleUpdate(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *RoleController) handleRemove(c *gin.Context) {
	req := &myRoleRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   true,
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
	id         rbac.RoleID
	pagination rbac.Pagination
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

	if inst.wantRequestID {
		str := c.Param("id")
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		inst.id = rbac.RoleID(n)
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
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
	ctx := inst.context
	ser := inst.controller.Service
	page := &inst.pagination
	q := &rbac.RoleQuery{}
	q.Pagination = *page
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Roles = list
	inst.body2.Pagination = page
	return nil
}

func (inst *myRoleRequest) doGetOne() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Roles = []*rbac.RoleDTO{o1}
	return nil
}

func (inst *myRoleRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := controllers.GetItemOnlyOne[rbac.RoleDTO](inst.body1.Roles)
	if err != nil {
		return err
	}
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Roles = []*rbac.RoleDTO{o2}
	return nil
}

func (inst *myRoleRequest) doUpdate() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := controllers.GetItemOnlyOne[rbac.RoleDTO](inst.body1.Roles)
	if err != nil {
		return err
	}
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.Roles = []*rbac.RoleDTO{o2}
	return nil
}

func (inst *myRoleRequest) doRemove() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	return ser.Delete(ctx, id)
}

////////////////////////////////////////////////////////////////////////////////
