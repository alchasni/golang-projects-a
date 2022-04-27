package domain

import (
	"reflect"

	"twatter/pkg/types"
)

type Role struct {
	ID   uint32
	Code types.Code
	Name string

	Permissions []Permission
}

func (r Role) Empty() bool {
	return reflect.DeepEqual(r, Role{})
}

func (r Role) PermissionCodes() []types.Code {
	permCodes := make([]types.Code, 0)
	for index, perm := range r.Permissions {
		permCodes[index] = perm.Code
	}

	return permCodes
}

type Roles struct {
	Items    []Role
	RowCount uint32
}
