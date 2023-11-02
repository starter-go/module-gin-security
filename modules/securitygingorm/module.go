package securitygingorm

import (
	"github.com/starter-go/application"
	securitygingorm "github.com/starter-go/security-gin-gorm"
	"github.com/starter-go/security-gin-gorm/gen/gen4ggsecurity"
	"github.com/starter-go/security-gin-gorm/gen/gen4ggsecuritytest"
)

// Module ...
func Module() application.Module {

	mb := securitygingorm.SrcMainModuleBuilder()
	mb.Components(gen4ggsecurity.ExportComponents)

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	mb := securitygingorm.SrcTestModuleBuilder()
	mb.Components(gen4ggsecuritytest.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
