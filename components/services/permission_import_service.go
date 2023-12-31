package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/vlog"
)

// PermissionImportService ...
type PermissionImportService interface {
	ImportFromResource(c context.Context) error
}

////////////////////////////////////////////////////////////////////////////////

// PermissionImportServiceImpl ...
type PermissionImportServiceImpl struct {

	//starter:component
	_as func(PermissionImportService) //starter:as("#")

	AC          application.Context    //starter:inject("context")
	PermService rbac.PermissionService //starter:inject("#")
	ResPath     string                 //starter:inject("${security.permissions.resource-file-path}")

}

func (inst *PermissionImportServiceImpl) _impl() PermissionImportService {
	return inst
}

// ImportFromResource ...
func (inst *PermissionImportServiceImpl) ImportFromResource(c context.Context) error {
	ser := inst.PermService
	src, err := inst.loadFromResource()
	if err != nil {
		return err
	}
	for _, perm1 := range src {
		perm2, err := ser.Insert(c, perm1)
		if err == nil {
			vlog.Info("import permission: %s %s", perm2.Method, perm2.Path)
		}
	}
	return nil
}

func (inst *PermissionImportServiceImpl) loadFromResource() ([]*rbac.PermissionDTO, error) {
	text, err := inst.AC.GetResources().ReadText(inst.ResPath)
	if err != nil {
		return nil, err
	}
	text = strings.ReplaceAll(text, "\r", "\n")
	rows := strings.Split(text, "\n")
	dst := make([]*rbac.PermissionDTO, 0)
	for _, row := range rows {
		perm, err := inst.parseRow(row)
		if err == nil {
			dst = append(dst, perm)
		}
	}
	return dst, nil
}

func (inst *PermissionImportServiceImpl) parseRow(row string) (*rbac.PermissionDTO, error) {
	const (
		prefix = "[GIN-debug]"
		ending = "-->"
	)
	// row format like:
	//[GIN-debug] POST   /api/dev.v1/Example       --> github.com/starter-go/...
	row = strings.TrimSpace(row)
	if !strings.HasPrefix(row, prefix) {
		return nil, fmt.Errorf("bad row: " + row)
	}
	iEnd := strings.Index(row, ending)
	if iEnd < 0 {
		return nil, fmt.Errorf("bad row: " + row)
	}
	methodAndPath := strings.TrimSpace(row[len(prefix):iEnd])
	iSpace := strings.IndexRune(methodAndPath, ' ')
	if iSpace < 0 {
		return nil, fmt.Errorf("bad row: " + row)
	}
	p1 := strings.TrimSpace(methodAndPath[0:iSpace])
	p2 := strings.TrimSpace(methodAndPath[iSpace+1:])
	perm := &rbac.PermissionDTO{
		Method: p1,
		Path:   p2,
	}
	return perm, nil
}
