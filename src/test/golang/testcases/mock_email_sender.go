package testcases

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/vlog"
)

// MockEmailSender ...
type MockEmailSender struct {

	//starter:component

	_as func(mails.DriverRegistry) //starter:as(".")

}

func (inst *MockEmailSender) _impl() mails.DriverRegistry {
	return inst
}

// ListRegistrations ...
func (inst *MockEmailSender) ListRegistrations() []*mails.DriverRegistration {
	r1 := &mails.DriverRegistration{
		Name:    "MockEmailSender",
		Enabled: true,
		Driver:  inst,
	}
	return []*mails.DriverRegistration{r1}
}

// Accept ...
func (inst *MockEmailSender) Accept(cfg *mails.Configuration) bool {
	return cfg.Driver == "mock-email"
}

// CreateDispatcher ...
func (inst *MockEmailSender) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	d := &myMockEmailSenderDispatcher{
		factory: inst,
		config:  cfg,
	}
	return d, nil
}

////////////////////////////////////////////////////////////////////////////////

type myMockEmailSenderDispatcher struct {
	factory *MockEmailSender
	config  *mails.Configuration
}

func (inst *myMockEmailSenderDispatcher) _impl() mails.Dispatcher {
	return inst
}

func (inst *myMockEmailSenderDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	a1 := msg.FromAddress
	a2 := inst.config.SenderAddress
	return a1 == a2
}

func (inst *myMockEmailSenderDispatcher) Send(c context.Context, msg *mails.Message) error {

	str := formatMailsMessage(msg)
	vlog.Info("myMockEmailSenderDispatcher.Send(email): %s", str)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func formatMailsMessage(msg *mails.Message) string {

	to := &strings.Builder{}
	for _, addr := range msg.ToAddresses {
		if to.Len() > 0 {
			to.WriteString(", ")
		}
		to.WriteString(addr.String())
	}

	text := msg.Content
	values := map[string]string{
		"title":        msg.Title,
		"content-type": msg.ContentType,
		"content":      string(text),
		"to":           to.String(),
	}

	bin, err := json.MarshalIndent(values, "", "\t")
	if err != nil {
		return err.Error()
	}
	str := string(bin)
	return str
}
