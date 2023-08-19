package gen4data

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForSecurityGinGormData ...
func ExportComForSecurityGinGormData(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
