package gen4test

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForSecurityGinGormTest ...
func ExportComForSecurityGinGormTest(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
