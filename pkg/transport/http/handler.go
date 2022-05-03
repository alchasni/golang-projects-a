package http

import (
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/core/service/comment"
	"golang-projects-a/pkg/core/service/organization"
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
