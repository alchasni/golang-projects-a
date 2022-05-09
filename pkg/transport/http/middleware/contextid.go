package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"golang-projects-a/pkg/contextid"
)

const (
	key_ContextID = "context_id"
)

func ContextIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := contextid.New(c.Request().Context())
			c.Set(key_ContextID, ctx)

			return next(c)
		}
	}
}

func ContextID(c echo.Context) context.Context {
	switch ctx := c.Get(key_ContextID).(type) {
	case context.Context:
		return ctx
	default:
		return contextid.New(c.Request().Context())
	}
}
