package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers/home"
)

////////////////////////////////////////////////////////////////////////////////

// APITester ...
type APITester struct {

	//starter:component()
	_as func(TestRunnable) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")

}

func (inst *APITester) _impl() TestRunnable {
	return inst
}

// GetRunnableInfo ...
func (inst *APITester) GetRunnableInfo() *TestRunnableInfo {
	info := &TestRunnableInfo{}
	info.Name = "APITester"
	info.Order = 0
	info.OnLoop = inst.run
	return info
}

func (inst *APITester) run() error {

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
