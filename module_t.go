package securitygingorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	moduleemail "github.com/starter-go/module-email"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/security-gin/modules/securitygin"
	"github.com/starter-go/security-gorm/modules/securitygorm"
)

const (
	theModuleName = "github.com/starter-go/security-gin-gorm"
	theModuleVer  = "v0.0.4"
	theModuleRev  = 4

	theSrcMainResPath = "src/main/resources"
	theSrcTestResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theSrcMainResFS embed.FS

//go:embed "src/test/resources"
var theSrcTestResFS embed.FS

// SrcMainModuleBuilder ...
func SrcMainModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theSrcMainResFS, theSrcMainResPath)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(securitygorm.Module())
	mb.Depend(securitygin.Module())
	mb.Depend(moduleemail.Module())

	return mb
}

// SrcTestModuleBuilder ...
func SrcTestModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theSrcTestResFS, theSrcTestResPath)

	mb.Depend(modulegormmysql.Module())
	mb.Depend(modulegormsqlserver.Module())

	return mb
}
