package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// SessionsCurrentTester ...
type SessionsCurrentTester struct {

	//starter:component()

	_as func(units.Units) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")

}

func (inst *SessionsCurrentTester) _impl() units.Units { return inst }

// Units  ...
func (inst *SessionsCurrentTester) Units(list []*units.Registration) []*units.Registration {

	info := &units.Registration{
		Name:     "SessionsCurrentTester",
		Test:     inst.run,
		Enabled:  true,
		Priority: 0 - OrderSessionsCurrent,
	}

	list = append(list, info)
	return list
}

func (inst *SessionsCurrentTester) run() error {

	req := &httpagent.Request{
		Method: http.MethodGet,
		URL:    "/api/v1/sessions/current",
	}

	client := inst.Agent.GetClient()
	resp, err := client.Execute(req)

	if err == nil {
		if resp.Status == http.StatusOK {
			inst.printSessionInfo(resp)
			return nil
		}
	}

	return inst.ResponseHandler.HandleResponse(resp, err)
}

func (inst *SessionsCurrentTester) printSessionInfo(resp *httpagent.Response) error {

	ent, err := resp.GetEntity()
	if err != nil {
		return err
	}

	text, err := ent.ReadText()
	if err != nil {
		return err
	}

	vlog.Info("session info text: %s", text)
	return nil
}
