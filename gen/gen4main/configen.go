package gen4main

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForGinSecurityMain ...
func ExportComForGinSecurityMain(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
