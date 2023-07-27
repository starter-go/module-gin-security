package moduleginsecurity

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modulegin"
	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/module-gin-security/gen/gen4main"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	"github.com/starter-go/security"
	"github.com/starter-go/starter"
)

const (
	theModuleName    = "github.com/starter-go/module-gin-security"
	theModuleVer     = "v0.0.0"
	theModuleRev     = 0
	theModuleResPath = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module ...
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4main.ExportComForGinSecurityMain)

	mb.Depend(starter.Module())
	mb.Depend(modulegin.Module())
	mb.Depend(security.Module())
	mb.Depend(modgorm.Module())
	mb.Depend(modulegormmysql.Module())

	return mb.Create()
}
