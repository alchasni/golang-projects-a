package http

import (
	"net/http"

	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/core/service/permission"
	"golang-projects-a/pkg/transport/http/middleware"
	"golang-projects-a/pkg/types"

	"github.com/labstack/echo/v4"
)

func (h HTTP) permissionFind(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Response struct {
		ID         uint32     `json:"id"`
		Code       types.Code `json:"code"`
		Name       string     `json:"name"`
		Restricted string     `json:"restricted"`
	}

	serviceResp, serviceErr := h.PermissionService.Find(ctx, permission.FindReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:         serviceResp.ID,
		Code:       serviceResp.Code,
		Name:       serviceResp.Name,
		Restricted: serviceResp.Restricted,
	})
}

func (h HTTP) permissionGetList(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Code       types.Code `query:"code"`
		Name       string     `query:"name"`
		Restricted string     `query:"restricted"`
		PageNo     int        `query:"page_no"`
		PageSize   int        `query:"page_size"`
	}
	type Permission struct {
		ID         uint32     `json:"id"`
		Code       types.Code `json:"code"`
		Name       string     `json:"name"`
		Restricted string     `json:"restricted"`
	}
	type Response struct {
		Items    []Permission `json:"items"`
		RowCount uint32       `json:"row_count"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.PermissionService.GetList(ctx, permission.GetListReq{
		Code:       req.Code,
		Name:       req.Name,
		Restricted: req.Restricted,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Items: func() []Permission {
			items := make([]Permission, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				items[index] = Permission{
					ID:         item.ID,
					Code:       item.Code,
					Name:       item.Name,
					Restricted: item.Restricted,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}

func (h HTTP) permissionCreate(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Code       types.Code `json:"code"`
		Name       string     `json:"name"`
		Restricted string     `json:"restricted"`
	}
	type Response struct {
		ID         uint32     `json:"id"`
		Code       types.Code `json:"code"`
		Name       string     `json:"name"`
		Restricted string     `json:"restricted"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.PermissionService.Create(ctx, permission.CreateReq{
		Code:       req.Code,
		Name:       req.Name,
		Restricted: req.Restricted,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:         serviceResp.ID,
		Code:       serviceResp.Code,
		Name:       serviceResp.Name,
		Restricted: serviceResp.Restricted,
	})
}

func (h HTTP) permissionUpdate(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Name       string `json:"name"`
		Restricted string `json:"restricted"`
	}
	type Response struct {
		ID         uint32     `json:"id"`
		Code       types.Code `json:"code"`
		Name       string     `json:"name"`
		Restricted string     `json:"restricted"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.PermissionService.Update(ctx, permission.UpdateReq{
		ID:         c.Param("id"),
		Name:       req.Name,
		Restricted: req.Restricted,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:         serviceResp.ID,
		Code:       serviceResp.Code,
		Name:       serviceResp.Name,
		Restricted: serviceResp.Restricted,
	})
}

func (h HTTP) permissionDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	serviceErr := h.PermissionService.Delete(ctx, permission.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
