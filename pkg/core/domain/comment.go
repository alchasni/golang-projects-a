package domain

import "time"

type Comment struct {
	ID             uint64 `json:"id,omitempty"`
	Comment        string `json:"comment" validate:"required"`
	OrganizationId uint64 `json:"organization_id" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Comments struct {
	Items    []Comment
	RowCount uint64
}
