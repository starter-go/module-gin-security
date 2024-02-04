package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
)

// GroupVO ...
type GroupVO struct {
	rbac.BaseVO

	Groups []*rbac.GroupDTO `json:"groups"`
}

////////////////////////////////////////////////////////////////////////////////

// GroupController ...
type GroupController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder  //starter:inject("#")
	Service   rbac.GroupService //starter:inject("#")

}

func (inst *GroupController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *GroupController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route:  inst.route,
		Groups: []string{"admin"},
	}
}

func (inst *GroupController) route(g libgin.RouterProxy) error {
	g = g.For("groups")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *GroupController) execute(req *myGroupRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *GroupController) handleGetList(c *gin.Context) {
	req := &myGroupRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: true,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetList)
}

func (inst *GroupController) handleGetOne(c *gin.Context) {
	req := &myGroupRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *GroupController) handleInsert(c *gin.Context) {
	req := &myGroupRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *GroupController) handleUpdate(c *gin.Context) {
	req := &myGroupRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *GroupController) handleRemove(c *gin.Context) {
	req := &myGroupRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myGroupRequest struct {
	// contexts
	controller *GroupController
	context    *gin.Context

	// flags
	wantRequestBody bool
	wantRequestID   bool
	wantRequestPage bool
	wantRequestRBAC bool

	// params
	id         rbac.GroupID
	pagination rbac.Pagination
	roles      rbac.RoleNameList

	// body
	body1 GroupVO
	body2 GroupVO
}

func (inst *myGroupRequest) open() error {

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
		inst.id = rbac.GroupID(n)
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
	}

	return nil
}

func (inst *myGroupRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myGroupRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	page := &inst.pagination
	q := &rbac.GroupQuery{}
	q.Pagination = *page
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Groups = list
	inst.body2.Pagination = page
	return nil
}

func (inst *myGroupRequest) doGetOne() error {
	return nil
}

func (inst *myGroupRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := controllers.GetItemOnlyOne[rbac.GroupDTO](inst.body1.Groups)
	if err != nil {
		return err
	}
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Groups = []*rbac.GroupDTO{o2}
	return nil
}

func (inst *myGroupRequest) doUpdate() error {
	return nil
}

func (inst *myGroupRequest) doRemove() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
