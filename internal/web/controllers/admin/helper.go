package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/module-security-gin-gorm/internal/web/controllers"
	"github.com/starter-go/security/rbac"
)

func getPagination(c *gin.Context) rbac.Pagination {
	return controllers.GetPagination(c)
}

func getItemOnlyOne[T any](list []*T) (*T, error) {
	return controllers.GetItemOnlyOne(list)
}
