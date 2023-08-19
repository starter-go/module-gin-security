package gen4core

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForSecurityGinGormCore ...
func ExportComForSecurityGinGormCore(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
