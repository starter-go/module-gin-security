package modulesecuritygingorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modulegin"
	"github.com/starter-go/libgorm/modgorm"
	moduleemail "github.com/starter-go/module-email"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/module-security-gin-gorm/gen/gen4core"
	"github.com/starter-go/module-security-gin-gorm/gen/gen4data"
	"github.com/starter-go/module-security-gin-gorm/gen/gen4web"

	"github.com/starter-go/security"
	"github.com/starter-go/starter"
)

const (
	theModuleName    = "github.com/starter-go/module-security-gin-gorm"
	theModuleVer     = "v0.0.1"
	theModuleRev     = 1
	theModuleResPath = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

func makeModule(nameSuffix string) *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + nameSuffix)
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(starter.Module())
	mb.Depend(security.Module())

	return mb
}

// ModuleData ...
func ModuleData() application.Module {
	mb := makeModule("#data")
	mb.Components(gen4data.ExportComForSecurityGinGormData)
	mb.Depend(modgorm.Module())
	mb.Depend(modulegormmysql.Module())
	mb.Depend(modulegormsqlserver.Module())
	return mb.Create()
}

// ModuleWeb ...
func ModuleWeb() application.Module {
	mb := makeModule("#web")
	mb.Components(gen4web.ExportComForSecurityGinGormWeb)
	mb.Depend(modulegin.Module())
	return mb.Create()
}

// ModuleCore ...
func ModuleCore() application.Module {
	mb := makeModule("#core")
	mb.Components(gen4core.ExportComForSecurityGinGormCore)
	return mb.Create()
}

// ModuleAll ...
func ModuleAll() application.Module {
	mb := makeModule("")
	mb.Components(func(r application.ComponentRegistry) error {
		return nil
	})
	mb.Depend(ModuleCore())
	mb.Depend(ModuleData())
	mb.Depend(ModuleWeb())
	mb.Depend(moduleemail.Module())
	return mb.Create()
}

// Module 这是 ModuleAll 的别名
func Module() application.Module {
	return ModuleAll()
}
