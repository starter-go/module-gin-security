package golang

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/vlog"
)

// Demo1vo ...
type Demo1vo struct {
	rbac.BaseVO

	Demo1 string `json:"demo1"`
}

////////////////////////////////////////////////////////////////////////////////

// Demo1Controller ...
type Demo1Controller struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	JWTser    jwt.Service      //starter:inject("#")

}

func (inst *Demo1Controller) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *Demo1Controller) Registration() *libgin.ControllerRegistration {
	// return &libgin.ControllerRegistration{Route: inst.route}
	return &libgin.ControllerRegistration{}
}

func (inst *Demo1Controller) route(g *gin.RouterGroup) error {
	g = g.Group("demo/jwt")

	g.GET("", inst.handleGet)
	g.GET(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.DELETE(":id", inst.handle)
	g.POST("", inst.handlePut)

	return nil
}

func (inst *Demo1Controller) handle(c *gin.Context) {
	req := &demo1request{
		controller: inst,
		context:    c,
	}
	err := req.open()
	if err == nil {
		err = req.do()
	}
	req.send(err)
}

func (inst *Demo1Controller) handleGet(c *gin.Context) {
	req := &demo1request{
		controller: inst,
		context:    c,
	}
	err := req.open()
	if err == nil {
		err = req.doGet()
	}
	req.send(err)
}

func (inst *Demo1Controller) handlePut(c *gin.Context) {
	req := &demo1request{
		controller: inst,
		context:    c,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type demo1request struct {
	controller *Demo1Controller
	context    *gin.Context

	wantRequestBody bool

	body1 Demo1vo
	body2 Demo1vo
}

func (inst *demo1request) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *demo1request) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *demo1request) do() error {
	return nil
}

func (inst *demo1request) doGet() error {

	ctx := inst.context
	ser := inst.controller.JWTser

	o, err := ser.GetDTO(ctx)
	if err == nil {
		j, err := json.Marshal(o)
		if err == nil {
			vlog.Info("jwt.json: %s", string(j))
		}
	}

	return err
}

func (inst *demo1request) doPut() error {

	now := lang.Now()

	o := &jwt.Token{
		ExpiredAt: now + 666000,
	}

	ctx := inst.context
	ser := inst.controller.JWTser
	return ser.SetDTO(ctx, o)
}

////////////////////////////////////////////////////////////////////////////////
