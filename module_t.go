package securitygingorm

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName = "github.com/starter-go/security-gin-gorm"
	theModuleVer  = "v1.0.55"
	theModuleRev  = 19
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

	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}
