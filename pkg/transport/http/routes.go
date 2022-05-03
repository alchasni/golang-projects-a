package http

import (
	"net/http"

	"golang-projects-a/pkg/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

const (
	pathKey_Ping = "ping"

	pathKey_CommentCreate  = "comment.create"
	pathKey_CommentDelete  = "comment.delete"
	pathKey_CommentFind    = "comment.find"
	pathKey_CommentUpdate  = "comment.update"
	pathKey_CommentGetList = "comment.get_list"

	pathKey_OrganizationCreate  = "organization.create"
	pathKey_OrganizationDelete  = "organization.delete"
	pathKey_OrganizationFind    = "organization.find"
	pathKey_OrganizationUpdate  = "organization.update"
	pathKey_OrganizationGetList = "organization.get_list"

	pathKey_UserCreate  = "user.create"
	pathKey_UserDelete  = "user.delete"
	pathKey_UserFind    = "user.find"
	pathKey_UserUpdate  = "user.update"
	pathKey_UserGetList = "user.get_list"
)

func (h HTTP) register() {
	contextIDMiddleware := middleware.ContextIDMiddleware()

	h.paths[pathKey_Ping] = h.e.GET("/ping", ping).Path

	api := h.e.Group("/api", contextIDMiddleware)
	apiV1 := api.Group("/v1")
	h.paths[pathKey_CommentCreate] = apiV1.POST("/comments", h.commentCreate).Path
	h.paths[pathKey_CommentDelete] = apiV1.DELETE("/comments/:id", h.commentDelete).Path
	h.paths[pathKey_CommentFind] = apiV1.GET("/comments/:id", h.commentFind).Path
	h.paths[pathKey_CommentUpdate] = apiV1.PUT("/comments/:id", h.commentUpdate).Path
	h.paths[pathKey_CommentGetList] = apiV1.GET("/comments", h.commentGetList).Path

	h.paths[pathKey_OrganizationCreate] = apiV1.POST("/orgs", h.organizationCreate).Path
	h.paths[pathKey_OrganizationDelete] = apiV1.DELETE("/orgs/:id", h.organizationDelete).Path
	h.paths[pathKey_OrganizationFind] = apiV1.GET("/orgs/:id", h.organizationFind).Path
	h.paths[pathKey_OrganizationUpdate] = apiV1.PUT("/orgs/:id", h.organizationUpdate).Path
	h.paths[pathKey_OrganizationGetList] = apiV1.GET("/orgs", h.organizationGetList).Path

	h.paths[pathKey_UserCreate] = apiV1.POST("/users", h.userCreate).Path
	h.paths[pathKey_UserDelete] = apiV1.DELETE("/users/:id", h.userDelete).Path
	h.paths[pathKey_UserFind] = apiV1.GET("/users/:id", h.userFind).Path
	h.paths[pathKey_UserUpdate] = apiV1.PUT("/users/:id", h.userUpdate).Path
	h.paths[pathKey_UserGetList] = apiV1.GET("/users", h.userGetList).Path

	//h.paths[pathKey_PermissionFind] = apiV1.GET("/permissions/:id", h.permissionFind).Path
	//h.paths[pathKey_PermissionGetList] = apiV1.GET("/permissions", h.permissionGetList).Path
	//h.paths[pathKey_PermissionCreate] = apiV1.POST("/permissions", h.permissionCreate).Path
	//h.paths[pathKey_PermissionUpdate] = apiV1.PUT("/permissions/:id", h.permissionUpdate).Path
	//h.paths[pathKey_PermissionDelete] = apiV1.DELETE("/permissions/:id", h.permissionDelete).Path
	//
	//h.paths[pathKey_RoleFind] = apiV1.GET("/roles/:id", h.roleFind).Path
	//h.paths[pathKey_RoleGetList] = apiV1.GET("/roles", h.roleGetList).Path
	//h.paths[pathKey_RoleCreate] = apiV1.POST("/roles", h.roleCreate).Path
	//h.paths[pathKey_RoleUpdate] = apiV1.PUT("/roles/:id", h.roleUpdate).Path
	//h.paths[pathKey_RoleDelete] = apiV1.DELETE("/roles/:id", h.roleDelete).Path
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"msg": "pong",
	})
}
