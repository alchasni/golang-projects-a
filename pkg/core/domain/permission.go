package domain

import "golang-projects-a/pkg/types"

type Permission struct {
	ID         uint32
	Code       types.Code
	Name       string
	Restricted string
}

func (p Permission) Empty() bool {
	return p == Permission{}
}

type Permissions struct {
	Items    []Permission
	RowCount uint32
}
