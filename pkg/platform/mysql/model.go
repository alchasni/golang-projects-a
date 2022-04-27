package mysql

import "twatter/pkg/types"

type Role struct {
	ID   uint32     `gorm:"primaryKey"`
	Code types.Code `gorm:"not null;size:50;uniqueIndex:uk_roles_code"`
	Name string     `gorm:"not null"`

	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	ID         uint32     `gorm:"primaryKey"`
	Code       types.Code `gorm:"not null;size:50;uniqueIndex:uk_permissions_code"`
	Name       string     `gorm:"not null"`
	Restricted string     `gorm:"not null;size:1;default:Y"`
}

type RolePermission struct {
	ID           uint32 `gorm:"primaryKey"`
	RoleID       uint32 `gorm:"not null;size:50;uniqueIndex:uk_role_permissions_role_id_permission_id"`
	PermissionID uint32 `gorm:"not null;size:50;uniqueIndex:uk_role_permissions_role_id_permission_id"`

	Role       *Role       `gorm:"foreignKey:RoleID;references:ID"`
	Permission *Permission `gorm:"foreignKey:PermissionID;references:ID"`
}
