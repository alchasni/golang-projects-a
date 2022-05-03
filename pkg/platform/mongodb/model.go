package mongodb

import (
	"time"
)

type Comment struct {
	ID             uint64 `json:"id" bson:"id,omitempty"`
	Comment        string `json:"comment" bson:"comment,omitempty"`
	OrganizationId uint64 `json:"organization_id" bson:"organization_id,omitempty"`

	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}

type Organization struct {
	ID   uint64 `json:"id" bson:"id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`

	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}

type User struct {
	ID             uint64 `json:"id" bson:"id,omitempty"`
	Username       string `json:"username" bson:"username,omitempty"`
	Password       string `json:"password" bson:"password,omitempty"`
	Email          string `json:"email" bson:"email,omitempty"`
	AvatarUrl      string `json:"avatar_url" bson:"avatar_url,omitempty"`
	OrganizationId uint64 `json:"organization_id" bson:"organization_id,omitempty"`
	FollowingCount uint64 `json:"following_count" bson:"following_count,omitempty"` // denormalized value. SOT: DB relation followers (out of scope)
	FollowerCount  uint64 `json:"follower_count" bson:"follower_count,omitempty"`   // denormalized value. SOT: DB relation followers (out of scope)

	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}
