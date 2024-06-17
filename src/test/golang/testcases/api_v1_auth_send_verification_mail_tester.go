package testcases

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers/home"
)

////////////////////////////////////////////////////////////////////////////////

// AuthSendVerificationMailTester ...
type AuthSendVerificationMailTester struct {

	//starter:component()
	_as func(TestRunnable) //starter:as(".")

	Agent           httpagent.Clients //starter:inject("#")
	ResponseHandler ResponseHandler   //starter:inject("#")

}

func (inst *AuthSendVerificationMailTester) _impl() TestRunnable {
	return inst
}

// GetRunnableInfo ...
func (inst *AuthSendVerificationMailTester) GetRunnableInfo() *TestRunnableInfo {
	info := &TestRunnableInfo{}
	info.Name = "AuthSendVerificationMailTester"
	info.Order = OrderAuthSendVerificationMail
	info.OnLoop = inst.run
	return info
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
