package http

import (
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/core/service/comment"
	"golang-projects-a/pkg/core/service/organization"
	"golang-projects-a/pkg/core/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang-projects-a/pkg/transport/http/middleware"
)

func (h HTTP) commentCreateByOrg(c echo.Context) error {
	ctx := middleware.ContextID(c)

	orgName := c.Param("name")
	type Request struct {
		Comment string `json:"comment"`
	}

	org, serviceErr := h.OrganizationService.FindByName(ctx, organization.FindByNameReq{Name: orgName})
	if serviceErr != nil {
		return serviceErr
	}
	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	_, serviceErr = h.CommentService.Create(ctx, comment.CreateReq{
		Comment:        req.Comment,
		OrganizationId: org.ID,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}

func (h HTTP) commentListByOrg(c echo.Context) error {
	ctx := middleware.ContextID(c)

	orgName := c.Param("name")

	org, serviceErr := h.OrganizationService.FindByName(ctx, organization.FindByNameReq{Name: orgName})
	if serviceErr != nil {
		return serviceErr
	}
	type Response struct {
		Comments []string `json:"comments"`
	}

	serviceResp, serviceErr := h.CommentService.GetList(ctx, comment.GetListReq{
		OrganizationId: org.ID,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Comments: func() []string {
			comments := make([]string, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				comments[index] = item.Comment
			}
			return comments
		}(),
	})
}

func (h HTTP) commentDeleteByOrg(c echo.Context) error {
	ctx := middleware.ContextID(c)

	orgName := c.Param("name")

	org, serviceErr := h.OrganizationService.FindByName(ctx, organization.FindByNameReq{Name: orgName})
	if serviceErr != nil {
		return serviceErr
	}

	serviceErr = h.CommentService.DeleteMany(ctx, comment.DeleteManyReq{
		OrganizationId: org.ID,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}

func (h HTTP) memberListByOrg(c echo.Context) error {
	ctx := middleware.ContextID(c)

	orgName := c.Param("name")
	type User struct {
		ID             uint64 `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		AvatarUrl      string `json:"avatar_url"`
		FollowingCount uint64 `json:"following_count"`
		FollowerCount  uint64 `json:"follower_count"`
	}
	type Response struct {
		Items    []User `json:"items"`
		RowCount uint64 `json:"row_count"`
	}

	org, serviceErr := h.OrganizationService.FindByName(ctx, organization.FindByNameReq{Name: orgName})
	if serviceErr != nil {
		return serviceErr
	}

	serviceResp, serviceErr := h.UserService.GetList(ctx, user.GetListReq{
		OrganizationId: org.ID,
		Limit:          100,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Items: func() []User {
			items := make([]User, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				items[index] = User{
					ID:             item.ID,
					Username:       item.Username,
					Email:          item.Email,
					AvatarUrl:      item.AvatarUrl,
					FollowingCount: item.FollowingCount,
					FollowerCount:  item.FollowerCount,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}
