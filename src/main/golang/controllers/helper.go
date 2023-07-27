package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/security/rbac"
)

func getItemOnlyOne[T any](list []*T) (*T, error) {

	if list == nil {
		return nil, fmt.Errorf("no item")
	}

	count := len(list)
	if count < 1 {
		return nil, fmt.Errorf("no item")
	} else if count > 1 {
		return nil, fmt.Errorf("there are items in list, more then one")
	}

	item := list[0]
	if item == nil {
		return nil, fmt.Errorf("no item")
	}

	return item, nil
}

func getPagination(c *gin.Context) rbac.Pagination {
	const (
		base    = 0
		bitSize = 0
	)
	pageNumb := c.Query("page")
	pageSize := c.Query("size")
	page, _ := strconv.ParseInt(pageNumb, base, bitSize)
	size, _ := strconv.ParseInt(pageSize, base, bitSize)

	if size < 1 {
		size = 5
	}
	if page < 1 {
		page = 1
	}

	return rbac.Pagination{
		Page: page,
		Size: int(size),
	}
}
