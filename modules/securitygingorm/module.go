package securitygingorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/httpagent/modules/httpagent"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	securitygingorm "github.com/starter-go/security-gin-gorm"
	"github.com/starter-go/security-gin-gorm/gen/main4ggsecurity"
	"github.com/starter-go/security-gin-gorm/gen/test4ggsecurity"
	"github.com/starter-go/security-gin/modules/securitygin"
	"github.com/starter-go/security-gorm/modules/securitygorm"
	"github.com/starter-go/security/modules/security"
)

// Module ...
func Module() application.Module {

	mb := securitygingorm.NewMainModule()
	mb.Components(main4ggsecurity.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(securitygorm.Module())
	mb.Depend(securitygin.Module())
	mb.Depend(security.Module())
	mb.Depend(mails.LibModule())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	mb := securitygingorm.NewTestModule()
	mb.Components(test4ggsecurity.ExportComponents)

	mb.Depend(Module())
	mb.Depend(httpagent.Module())
	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}
