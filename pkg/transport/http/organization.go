package http

import (
	"golang-projects-a/pkg/core/service/organization"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/transport/http/middleware"
)

func (h HTTP) organizationFind(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Response struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	serviceResp, serviceErr := h.OrganizationService.Find(ctx, organization.FindReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Name: serviceResp.Name,
	})
}

func (h HTTP) organizationGetList(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		ID   uint64 `query:"id"`
		Name string `query:"name"`

		Limit  int `query:"limit"`
		Offset int `query:"offset"`
	}
	type Organization struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	type Response struct {
		Items    []Organization `json:"items"`
		RowCount uint64         `json:"row_count"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.OrganizationService.GetList(ctx, organization.GetListReq{
		ID:     req.ID,
		Name:   req.Name,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Items: func() []Organization {
			items := make([]Organization, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				items[index] = Organization{
					ID:   item.ID,
					Name: item.Name,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}

func (h HTTP) organizationCreate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Name string `json:"name"`
	}
	type Response struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.OrganizationService.Create(ctx, organization.CreateReq{
		Name: req.Name,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Name: serviceResp.Name,
	})
}

func (h HTTP) organizationUpdate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Name string `json:"name"`
	}
	type Response struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.OrganizationService.Update(ctx, organization.UpdateReq{
		ID:   c.Param("id"),
		Name: req.Name,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Name: serviceResp.Name,
	})
}

func (h HTTP) organizationDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)

	serviceErr := h.OrganizationService.Delete(ctx, organization.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
