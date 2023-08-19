package gen4web

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForSecurityGinGormWeb ...
func ExportComForSecurityGinGormWeb(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
