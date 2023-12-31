package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/services"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
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

	Responder     libgin.Responder                 //starter:inject("#")
	Service       rbac.PermissionService           //starter:inject("#")
	ImportService services.PermissionImportService //starter:inject("#")

}

func (inst *PermissionController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *PermissionController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route:  inst.route,
		Groups: []string{"admin"},
	}
}

func (inst *PermissionController) route(g libgin.RouterProxy) error {
	g = g.For("permissions")

	g.GET("", inst.handleGetList)
	g.GET(":id", inst.handleGetOne)

	g.PUT(":id", inst.handleUpdate)
	g.DELETE(":id", inst.handleRemove)

	g.POST("", inst.handleInsert)
	g.POST("apply", inst.handleApply)
	g.POST("import-from-resource", inst.handleImportFromResource)

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
		wantRequestPage: true,
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
		wantRequestID:   true,
	}
	inst.execute(req, req.doGetOne)
}

func (inst *PermissionController) handleInsert(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doInsert)
}

func (inst *PermissionController) handleApply(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doApply)
}

func (inst *PermissionController) handleImportFromResource(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   false,
	}
	inst.execute(req, req.doImportFromResource)
}

func (inst *PermissionController) handleUpdate(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
		wantRequestPage: false,
		wantRequestID:   true,
	}
	inst.execute(req, req.doUpdate)
}

func (inst *PermissionController) handleRemove(c *gin.Context) {
	req := &myPermissionRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: false,
		wantRequestPage: false,
		wantRequestID:   true,
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
	id         rbac.PermissionID
	pagination rbac.Pagination
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

	if inst.wantRequestID {
		str := c.Param("id")
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		inst.id = rbac.PermissionID(n)
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
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
	ctx := inst.context
	ser := inst.controller.Service
	page := &rbac.Pagination{}
	q := &rbac.PermissionQuery{}
	q.Pagination = inst.pagination
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	*page = q.Pagination
	inst.body2.Permissions = list
	inst.body2.Pagination = page
	return nil
}

func (inst *myPermissionRequest) doGetOne() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Permissions = []*rbac.PermissionDTO{o1}
	return nil
}

func (inst *myPermissionRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	o1, err := controllers.GetItemOnlyOne[rbac.PermissionDTO](inst.body1.Permissions)
	if err != nil {
		return err
	}
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Permissions = []*rbac.PermissionDTO{o2}
	return nil
}

func (inst *myPermissionRequest) doApply() error {
	// ctx := inst.context
	ser := inst.controller.Service
	ser.GetCache().Clear()
	return nil
}

func (inst *myPermissionRequest) doImportFromResource() error {
	ctx := inst.context
	ser := inst.controller.ImportService
	return ser.ImportFromResource(ctx)
}

func (inst *myPermissionRequest) doUpdate() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	o1, err := controllers.GetItemOnlyOne[rbac.PermissionDTO](inst.body1.Permissions)
	if err != nil {
		return err
	}
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.Permissions = []*rbac.PermissionDTO{o2}
	return nil
}

func (inst *myPermissionRequest) doRemove() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	return ser.Delete(ctx, id)
}

////////////////////////////////////////////////////////////////////////////////
