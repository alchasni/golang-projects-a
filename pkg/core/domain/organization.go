package domain

import "time"

type Organization struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Organizations struct {
	Items    []Organization
	RowCount uint64
}
