package securitygingorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security-gin/modules/securitygin"
	"github.com/starter-go/security-gorm/modules/securitygorm"
	"github.com/starter-go/security/modules/security"

	moduleemail "github.com/starter-go/module-email"
)

const (
	theModuleName = "github.com/starter-go/security-gin-gorm"
	theModuleVer  = "v1.0.48"
	theModuleRev  = 15
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule ...
func NewMainModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(securitygorm.Module())
	mb.Depend(securitygin.Module())
	mb.Depend(security.Module())
	mb.Depend(moduleemail.Module())

	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb
}
