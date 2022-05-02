package domain

import "time"

type User struct {
	ID             uint64 `json:"id,omitempty"`
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Email          string `json:"email" validate:"required"`
	AvatarUrl      string `json:"avatar_url" validate:"required"`
	OrganizationId uint64 `json:"organization_id"`
	FollowingCount uint64 `json:"following_count"` // denormalized value. SOT: DB relation followers (out of scope)
	FollowerCount  uint64 `json:"follower_count"`  // denormalized value. SOT: DB relation followers (out of scope)

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Users struct {
	Items    []User
	RowCount uint64
}
