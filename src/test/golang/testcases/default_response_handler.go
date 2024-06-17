package testcases

import (
	"fmt"
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/vlog"
)

// ResponseHandler ...
type ResponseHandler interface {
	HandleResponse(resp *httpagent.Response, err error) error
}

////////////////////////////////////////////////////////////////////////////////

// DefaultResponseHandler ...
type DefaultResponseHandler struct {

	//starter:component

	_as func(ResponseHandler) //starter:as("#")
}

func (inst *DefaultResponseHandler) _impl() ResponseHandler {
	return inst
}

// HandleResponse ...
func (inst *DefaultResponseHandler) HandleResponse(resp *httpagent.Response, err error) error {

	if resp == nil {
		return nil
	}

	ent, err := resp.GetEntity()
	if err != nil {
		return err
	}

	txt, err := ent.ReadText()
	if err != nil {
		return err
	}

	code := resp.Status
	msg := resp.Message
	if code == http.StatusOK {
		return nil
	}

	// error
	vlog.Error("HTTP %s", msg)
	vlog.Error("HTTP error detail: %s", txt)
	return fmt.Errorf("HTTP %s", msg)
}
