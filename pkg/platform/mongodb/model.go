package mongodb

import (
	"time"
)

type Organization struct {
	ID   uint64 `json:"id" bson:"id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`

	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}

type User struct {
	ID        uint64 `json:"id" bson:"id,omitempty"`
	Username  string `json:"username" bson:"username,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	AvatarUrl string `json:"avatar_url" bson:"avatar_url,omitempty"`

	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deleted_at,omitempty"`
}
