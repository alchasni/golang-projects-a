package http

import (
	"net/http"

	"golang-projects-a/pkg/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

const (
	pathKey_Ping          = "ping"
	pathKey_HandlerCreate = "handler.create"
	pathKey_HandlerList   = "handler.list"
	pathKey_HandlerDelete = "handler.delete"
	pathKey_HandlerMember = "handler.member"

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

	// main endpoint
	h.paths[pathKey_HandlerCreate] = apiV1.POST("/orgs/:name/comments", h.commentCreateByOrg).Path
	h.paths[pathKey_HandlerList] = apiV1.GET("/orgs/:name/comments", h.commentListByOrg).Path
	h.paths[pathKey_HandlerDelete] = apiV1.DELETE("/orgs/:name/comments", h.commentDeleteByOrg).Path
	h.paths[pathKey_HandlerMember] = apiV1.GET("/orgs/:name/members", h.memberListByOrg).Path

	// helper endpoint for populate and checking data
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
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"msg": "pong",
	})
}
