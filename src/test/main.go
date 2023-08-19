package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	modulesecuritygingorm "github.com/starter-go/module-security-gin-gorm"
	"github.com/starter-go/module-security-gin-gorm/gen/gen4test"

	"github.com/starter-go/starter"
)

const (
	theModuleResPath = "resources"
)

//go:embed "resources"
var theModuleResFS embed.FS

// module ...
func module() application.Module {

	parent := modulesecuritygingorm.ModuleAll()

	mb := &application.ModuleBuilder{}
	mb.Name(parent.Name() + "#test")
	mb.Version(parent.Version())
	mb.Revision(parent.Revision())
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4test.ExportComForSecurityGinGormTest)

	mb.Depend(parent)
	mb.Depend(modulegormmysql.Module())

	return mb.Create()
}

func main() {
	i := starter.Init(os.Args)
	i.MainModule(module())
	i.WithPanic(true).Run()
}
