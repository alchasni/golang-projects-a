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
