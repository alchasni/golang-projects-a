package domain

type RolePermission struct {
	ID           uint32
	RoleID       uint32
	PermissionID uint32

	Role       *Role
	Permission *Permission
}

func (r RolePermission) Empty() bool {
	return r == RolePermission{}
}

type RolePermissions struct {
	Items    []RolePermission
	RowCount uint32
}
