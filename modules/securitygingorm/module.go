package securitygingorm

import (
	"github.com/starter-go/application"
	securitygingorm "github.com/starter-go/security-gin-gorm"
	"github.com/starter-go/security-gin-gorm/gen/main4ggsecurity"
	"github.com/starter-go/security-gin-gorm/gen/test4ggsecurity"
)

// Module ...
func Module() application.Module {

	mb := securitygingorm.NewMainModule()
	mb.Components(main4ggsecurity.ExportComponents)

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	mb := securitygingorm.NewTestModule()
	mb.Components(test4ggsecurity.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
