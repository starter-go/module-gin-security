package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers/home"
	"github.com/starter-go/units"
)

// AuthSignUpTester ...
type AuthSignUpTester struct {

	//starter:component()

	_as func(units.Units) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")
}

func (inst *AuthSignUpTester) _impl() units.Units { return inst }

// Units  ...
func (inst *AuthSignUpTester) Units(list []*units.Registration) []*units.Registration {

	info := &units.Registration{
		Name:     "AuthSignUpTester",
		Test:     inst.run,
		Enabled:  true,
		Priority: 0 - OrderAuthSignUp,
	}

	list = append(list, info)
	return list
}

func (inst *AuthSignUpTester) run() error {

	req := &httpagent.Request{
		Method: http.MethodPost,
		URL:    "/api/v1/auth/sign-in",
	}

	a1 := &rbac.AuthDTO{
		Mechanism: "password",
		Account:   "demo",
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
