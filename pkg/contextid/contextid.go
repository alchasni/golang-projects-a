package contextid

import (
	"context"

	"github.com/google/uuid"
)

type key string

const key_ContextID key = "context_id"

func New(ctx context.Context) context.Context {
	return NewWithValue(ctx, uuid.New().String())
}

func NewWithValue(ctx context.Context, value string) context.Context {
	if value == "" {
		value = uuid.New().String()
	}
	return context.WithValue(ctx, key_ContextID, value)
}

func Value(ctx context.Context) string {
	ctxID, ok := ctx.Value(key_ContextID).(string)
	if !ok {
		return ""
	}

	return ctxID
}
