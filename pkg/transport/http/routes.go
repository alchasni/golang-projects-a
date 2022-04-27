package http

import (
	"net/http"

	"twatter/pkg/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

const (
	pathKey_Ping = "ping"

	pathKey_PermissionFind    = "permission.find"
	pathKey_PermissionGetList = "permission.get_list"
	pathKey_PermissionCreate  = "permission.create"
	pathKey_PermissionUpdate  = "permission.update"
	pathKey_PermissionDelete  = "permission.delete"

	pathKey_RoleFind    = "role.find"
	pathKey_RoleGetList = "role.get_list"
	pathKey_RoleCreate  = "role.create"
	pathKey_RoleUpdate  = "role.update"
	pathKey_RoleDelete  = "role.delete"
)

func (h HTTP) register() {
	contextIDMiddleware := middleware.ContextIDMiddleware()

	h.paths[pathKey_Ping] = h.e.GET("/ping", ping).Path

	api := h.e.Group("/api", contextIDMiddleware)
	apiV1 := api.Group("/v1")

	h.paths[pathKey_PermissionFind] = apiV1.GET("/permissions/:id", h.permissionFind).Path
	h.paths[pathKey_PermissionGetList] = apiV1.GET("/permissions", h.permissionGetList).Path
	h.paths[pathKey_PermissionCreate] = apiV1.POST("/permissions", h.permissionCreate).Path
	h.paths[pathKey_PermissionUpdate] = apiV1.PUT("/permissions/:id", h.permissionUpdate).Path
	h.paths[pathKey_PermissionDelete] = apiV1.DELETE("/permissions/:id", h.permissionDelete).Path

	h.paths[pathKey_RoleFind] = apiV1.GET("/roles/:id", h.roleFind).Path
	h.paths[pathKey_RoleGetList] = apiV1.GET("/roles", h.roleGetList).Path
	h.paths[pathKey_RoleCreate] = apiV1.POST("/roles", h.roleCreate).Path
	h.paths[pathKey_RoleUpdate] = apiV1.PUT("/roles/:id", h.roleUpdate).Path
	h.paths[pathKey_RoleDelete] = apiV1.DELETE("/roles/:id", h.roleDelete).Path
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"msg": "pong",
	})
}
