package testcases

import (
	"context"

	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/vlog"
)

// MockSMSSender ...
type MockSMSSender struct {

	//starter:component

	_as func(mails.DriverRegistry) //starter:as(".")

}

func (inst *MockSMSSender) _impl() mails.DriverRegistry {
	return inst
}

// ListRegistrations ...
func (inst *MockSMSSender) ListRegistrations() []*mails.DriverRegistration {
	r1 := &mails.DriverRegistration{
		Name:    "MockSMSSender",
		Enabled: true,
		Driver:  inst,
	}
	return []*mails.DriverRegistration{r1}
}

// Accept ...
func (inst *MockSMSSender) Accept(cfg *mails.Configuration) bool {
	return cfg.Driver == "mock-sms"
}

// CreateDispatcher ...
func (inst *MockSMSSender) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	d := &myMockSMSSenderDispatcher{
		factory: inst,
		config:  cfg,
	}
	return d, nil
}

////////////////////////////////////////////////////////////////////////////////

type myMockSMSSenderDispatcher struct {
	factory *MockSMSSender
	config  *mails.Configuration
}

func (inst *myMockSMSSenderDispatcher) _impl() mails.Dispatcher {
	return inst
}

func (inst *myMockSMSSenderDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	a1 := msg.FromAddress
	a2 := inst.config.SenderAddress
	return a1 == a2
}

func (inst *myMockSMSSenderDispatcher) Send(c context.Context, msg *mails.Message) error {

	str := formatMailsMessage(msg)
	vlog.Info("myMockSMSSenderDispatcher.Send(sms): %s", str)

	return nil
}

////////////////////////////////////////////////////////////////////////////////
