package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	moduleginsecurity "github.com/starter-go/module-gin-security"
	"github.com/starter-go/module-gin-security/gen/gen4test"
	"github.com/starter-go/starter"
)

const (
	theModuleResPath = "resources"
)

//go:embed "resources"
var theModuleResFS embed.FS

// module ...
func module() application.Module {

	parent := moduleginsecurity.Module()

	mb := &application.ModuleBuilder{}
	mb.Name(parent.Name() + "#test")
	mb.Version(parent.Version())
	mb.Revision(parent.Revision())
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4test.ExportComForGinSecurityTest)

	mb.Depend(parent)

	return mb.Create()
}

func main() {
	i := starter.Init(os.Args)
	i.MainModule(module())
	i.WithPanic(true).Run()
}
