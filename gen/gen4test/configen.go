package gen4test

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForGinSecurityTest ...
func ExportComForGinSecurityTest(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
