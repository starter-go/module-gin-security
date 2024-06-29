package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
)

// RegionVO ...
type RegionVO struct {
	rbac.BaseVO

	Regions []*rbac.RegionDTO `json:"regions"`
}

////////////////////////////////////////////////////////////////////////////////

// RegionController ...
type RegionController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder   //starter:inject("#")
	Service   rbac.RegionService //starter:inject("#")

}

func (inst *RegionController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *RegionController) Registration() *libgin.ControllerRegistration {
	apiGroups := myAPIGroup()
	return &libgin.ControllerRegistration{
		Route:  inst.route,
		Groups: apiGroups,
	}
}

func (inst *RegionController) route(g libgin.RouterProxy) error {
	g = g.For("regions")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)
	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)
	g.POST("", inst.handleInsert)

	return nil
}

func (inst *RegionController) execute(req *myRegionRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *RegionController) handleGetList(c *gin.Context) {
	req := &myRegionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestID:   false,
		wantRequestPage: true,
	}
	inst.execute(req, req.doGetList)
}

func (inst *RegionController) handleGetOne(c *gin.Context) {
	req := &myRegionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *RegionController) handleInsert(c *gin.Context) {
	req := &myRegionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *RegionController) handleUpdate(c *gin.Context) {
	req := &myRegionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *RegionController) handleRemove(c *gin.Context) {
	req := &myRegionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myRegionRequest struct {
	// contexts
	controller *RegionController
	context    *gin.Context

	// flags
	wantRequestBody bool
	wantRequestID   bool
	wantRequestPage bool
	wantRequestRBAC bool

	// params
	id         rbac.RegionID
	pagination rbac.Pagination
	roles      rbac.RoleNameList

	// body
	body1 RegionVO
	body2 RegionVO
}

func (inst *myRegionRequest) open() error {

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
		inst.id = rbac.RegionID(n)
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
	}

	return nil
}

func (inst *myRegionRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myRegionRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	page := &inst.pagination
	q := &rbac.RegionQuery{}
	q.Pagination = *page
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Regions = list
	inst.body2.Pagination = page
	return nil
}

func (inst *myRegionRequest) doGetOne() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Regions = []*rbac.RegionDTO{o1}
	return nil
}

func (inst *myRegionRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := controllers.GetItemOnlyOne[rbac.RegionDTO](inst.body1.Regions)
	if err != nil {
		return err
	}
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Regions = []*rbac.RegionDTO{o2}
	return nil
}

func (inst *myRegionRequest) doUpdate() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := controllers.GetItemOnlyOne[rbac.RegionDTO](inst.body1.Regions)
	if err != nil {
		return err
	}
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.Regions = []*rbac.RegionDTO{o2}
	return nil
}

func (inst *myRegionRequest) doRemove() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	return ser.Delete(ctx, id)
}

////////////////////////////////////////////////////////////////////////////////
