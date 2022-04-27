package domain

import (
	"time"

	"golang-projects-a/pkg/types"
)

type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	RoleCode  types.Code

	Role *Role
}

type Users struct {
	Items    []User
	RowCount uint64
}
