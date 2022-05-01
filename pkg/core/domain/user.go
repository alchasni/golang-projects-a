package domain

import "time"

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
	AvatarUrl string `json:"avatar_url" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	//OrganisationId string
	//FollowingCount uint64 // denormalized value. SOT: DB relation followers (out of scope)
	//FollowerCount  uint64 // denormalized value. SOT: DB relation followers (out of scope)
}

type Users struct {
	Items    []User
	RowCount uint64
}
