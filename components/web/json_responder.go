package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/rbac"
)

// JSONGinResponder 用来发送 JSON 形式的 HTTP 响应
type JSONGinResponder struct {

	//--starter:component
	_as func(libgin.Responder) //starter:as(".")
}

func (inst *JSONGinResponder) _impl() (libgin.ResponderRegistry, libgin.Responder) {
	return inst, inst
}

// ListRegistrations ...
func (inst *JSONGinResponder) ListRegistrations() []*libgin.ResponderRegistration {
	r1 := &libgin.ResponderRegistration{
		Name:      "JSON",
		Enabled:   true,
		Responder: inst,
		Priority:  100,
	}
	return []*libgin.ResponderRegistration{r1}
}

// Accept ...
func (inst *JSONGinResponder) Accept(resp *libgin.Response) bool {
	b, err := inst.getBaseVO(resp)
	return (err == nil) && (b != nil)
}

// Send ...
func (inst *JSONGinResponder) Send(resp *libgin.Response) {

	b, err := inst.getBaseVO(resp)
	if err != nil {
		panic(err)
	}

	ctx := resp.Context
	data := resp.Data
	now := time.Now()
	err = resp.Error

	status, errStr := inst.getStatusAndError(b, resp)

	b.Error = errStr
	b.Status = status
	b.Time = now
	b.Timestamp = lang.NewTime(now)

	ctx.JSON(status, data)
}

func (inst *JSONGinResponder) getStatusAndError(b *rbac.BaseVO, r *libgin.Response) (int, string) {

	status := b.Status
	err1 := r.Error
	err2 := b.Error

	if err1 != nil {
		webErr, ok := err1.(web.Error)
		if ok {
			status = webErr.Status()
		} else {
			status = http.StatusInternalServerError
		}
		err2 = err1.Error()
	}

	if status == 0 && err1 == nil {
		status = http.StatusOK
	}

	if b.Message == "" {
		b.Message = http.StatusText(status)
	}

	return status, err2
}

func (inst *JSONGinResponder) getBaseVO(resp *libgin.Response) (*rbac.BaseVO, error) {
	const err1 = "cannot get BaseVO from rbac.VOGetter"
	getter, ok := resp.Data.(rbac.VOGetter)
	if !ok || getter == nil {
		return nil, fmt.Errorf(err1)
	}
	o := getter.GetVO()
	if o == nil {
		return nil, fmt.Errorf(err1)
	}
	return o, nil
}
