package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers/home"
	"github.com/starter-go/units"
)

////////////////////////////////////////////////////////////////////////////////

// AuthSendVerificationMailTester ...
type AuthSendVerificationMailTester struct {

	//starter:component()

	_as func(units.Units) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")

}

func (inst *AuthSendVerificationMailTester) _impl() units.Units { return inst }

// Units  ...
func (inst *AuthSendVerificationMailTester) Units(list []*units.Registration) []*units.Registration {

	info := &units.Registration{
		Name:     "AuthSendVerificationMailTester",
		Test:     inst.run,
		Enabled:  true,
		Priority: 0 - OrderAuthSendVerificationMail,
	}

	list = append(list, info)
	return list
}

func (inst *AuthSendVerificationMailTester) run() error {

	req := &httpagent.Request{
		Method: http.MethodPost,
		URL:    "/api/v1/auth/send-verification-mail",
	}

	a1 := &rbac.AuthDTO{
		Mechanism: "email",
		Account:   "test@example.com",
		Secret:    "demo",
	}
	body1 := &home.AuthVO{}
	body1.Auth = []*rbac.AuthDTO{a1}

	ent1 := httpagent.NewEntityWithJSON(body1)
	req.SetEntity(ent1)

	client := inst.Agent.GetClient()
	resp, err := client.Execute(req)

	return inst.ResponseHandler.HandleResponse(resp, err)
}
