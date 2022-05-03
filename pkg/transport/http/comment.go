package http

import (
	"golang-projects-a/pkg/core/service/comment"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/transport/http/middleware"
)

func (h HTTP) commentFind(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Response struct {
		ID             uint64 `json:"id"`
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}

	serviceResp, serviceErr := h.CommentService.Find(ctx, comment.FindReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:             serviceResp.ID,
		Comment:        serviceResp.Comment,
		OrganizationId: serviceResp.OrganizationId,
	})
}

func (h HTTP) commentGetList(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		ID             uint64 `query:"id"`
		OrganizationId uint64 `query:"organization_id"`

		Limit  int `query:"limit"`
		Offset int `query:"offset"`
	}
	type User struct {
		ID             uint64 `json:"id"`
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}
	type Response struct {
		Items    []User `json:"items"`
		RowCount uint64 `json:"row_count"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.CommentService.GetList(ctx, comment.GetListReq{
		ID:             req.ID,
		OrganizationId: req.OrganizationId,
		Limit:          req.Limit,
		Offset:         req.Offset,
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
					Comment:        item.Comment,
					OrganizationId: item.OrganizationId,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}

func (h HTTP) commentCreate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}
	type Response struct {
		ID             uint64 `json:"id"`
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.CommentService.Create(ctx, comment.CreateReq{
		Comment:        req.Comment,
		OrganizationId: req.OrganizationId,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:             serviceResp.ID,
		Comment:        serviceResp.Comment,
		OrganizationId: serviceResp.OrganizationId,
	})
}

func (h HTTP) commentUpdate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}
	type Response struct {
		ID             uint64 `json:"id"`
		Comment        string `json:"comment"`
		OrganizationId uint64 `json:"organization_id"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.CommentService.Update(ctx, comment.UpdateReq{
		ID:             c.Param("id"),
		Comment:        req.Comment,
		OrganizationId: req.OrganizationId,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:             serviceResp.ID,
		Comment:        serviceResp.Comment,
		OrganizationId: serviceResp.OrganizationId,
	})
}

func (h HTTP) commentDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)

	serviceErr := h.CommentService.Delete(ctx, comment.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
