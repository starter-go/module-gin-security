package testcases

import (
	"github.com/starter-go/httpagent"
)

// MyHTTPAgentFilter ...
type MyHTTPAgentFilter struct {

	//starter:component()
	_as func(httpagent.FilterRegistry) //starter:as(".")

}

func (inst *MyHTTPAgentFilter) _impl() (httpagent.Filter, httpagent.FilterRegistry) {
	return inst, inst
}

// FilterRegistrations ...
func (inst *MyHTTPAgentFilter) FilterRegistrations() []*httpagent.FilterRegistration {

	r1 := &httpagent.FilterRegistration{
		Name:    "",
		Enabled: true,
		Order:   0,
		Filter:  inst,
	}

	return []*httpagent.FilterRegistration{r1}
}

// Handle ...
func (inst *MyHTTPAgentFilter) Handle(c *httpagent.Context, chain httpagent.FilterChain) error {

	return chain.Handle(c)
}
