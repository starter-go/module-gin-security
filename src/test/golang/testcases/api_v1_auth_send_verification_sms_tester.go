package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers/home"
	"github.com/starter-go/units"
)

////////////////////////////////////////////////////////////////////////////////

// AuthSendVerificationSMSTester ...
type AuthSendVerificationSMSTester struct {

	//starter:component()

	_as func(units.Units) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")

}

func (inst *AuthSendVerificationSMSTester) _impl() units.Units { return inst }

// Units  ...
func (inst *AuthSendVerificationSMSTester) Units(list []*units.Registration) []*units.Registration {

	info := &units.Registration{
		Name:     "AuthSendVerificationSMSTester",
		Test:     inst.run,
		Enabled:  true,
		Priority: 0 - OrderAuthSendVerificationSMS,
	}

	list = append(list, info)
	return list
}

func (inst *AuthSendVerificationSMSTester) run() error {

	req := &httpagent.Request{
		Method: http.MethodPost,
		URL:    "/api/v1/auth/send-verification-sms",
	}

	a1 := &rbac.AuthDTO{
		Mechanism: "sms",
		Account:   "+86-13366669999",
		Secret:    "n/a",
	}
	body1 := &home.AuthVO{}
	body1.Auth = []*rbac.AuthDTO{a1}
	ent1 := httpagent.NewEntityWithJSON(body1)
	req.SetEntity(ent1)

	client := inst.Agent.GetClient()
	resp, err := client.Execute(req)

	return inst.ResponseHandler.HandleResponse(resp, err)
}
