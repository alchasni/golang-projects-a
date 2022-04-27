package http

import (
	"net/http"

	"twatter/pkg/core/service"
	"twatter/pkg/core/service/role"
	"twatter/pkg/transport/http/middleware"
	"twatter/pkg/types"

	"github.com/labstack/echo/v4"
)

func (h HTTP) roleFind(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Permission struct {
		ID   uint32     `json:"id"`
		Code types.Code `json:"code"`
		Name string     `json:"name"`
	}
	type Response struct {
		ID          uint32       `json:"id"`
		Code        types.Code   `json:"code"`
		Name        string       `json:"name"`
		Permissions []Permission `json:"permissions"`
	}

	serviceResp, serviceErr := h.RoleService.Find(ctx, role.FindReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Code: serviceResp.Code,
		Name: serviceResp.Name,
		Permissions: func() []Permission {
			perms := make([]Permission, len(serviceResp.Permissions))
			for index, perm := range serviceResp.Permissions {
				perms[index] = Permission{
					ID:   perm.ID,
					Code: perm.Code,
					Name: perm.Name,
				}
			}
			return perms
		}(),
	})
}

func (h HTTP) roleGetList(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Code            types.Code   `query:"code"`
		Name            string       `query:"name"`
		PermissionCodes []types.Code `query:"permission_code"`
		PageNo          int          `query:"page_no"`
		PageSize        int          `query:"page_size"`
	}
	type Role struct {
		ID   uint32     `json:"id"`
		Code types.Code `json:"code"`
		Name string     `json:"name"`
	}
	type Response struct {
		Items    []Role `json:"items"`
		RowCount uint32 `json:"row_count"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.RoleService.GetList(ctx, role.GetListReq{
		Code:            req.Code,
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Items: func() []Role {
			items := make([]Role, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				items[index] = Role{
					ID:   item.ID,
					Code: item.Code,
					Name: item.Name,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}

func (h HTTP) roleCreate(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Code            types.Code   `json:"code"`
		Name            string       `json:"name"`
		PermissionCodes []types.Code `json:"permission_codes"`
	}
	type Permission struct {
		ID   uint32     `json:"id"`
		Code types.Code `json:"code"`
		Name string     `json:"name"`
	}
	type Response struct {
		ID          uint32       `json:"id"`
		Code        types.Code   `json:"code"`
		Name        string       `json:"name"`
		Permissions []Permission `json:"permissions"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.RoleService.Create(ctx, role.CreateReq{
		Code:            req.Code,
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Code: serviceResp.Code,
		Name: serviceResp.Name,
		Permissions: func() []Permission {
			perms := make([]Permission, len(serviceResp.Permissions))
			for index, perm := range serviceResp.Permissions {
				perms[index] = Permission{
					ID:   perm.ID,
					Code: perm.Code,
					Name: perm.Name,
				}
			}
			return perms
		}(),
	})
}

func (h HTTP) roleUpdate(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	type Request struct {
		Name            string       `json:"name"`
		PermissionCodes []types.Code `json:"permission_codes"`
	}
	type Permission struct {
		ID   uint32     `json:"id"`
		Code types.Code `json:"code"`
		Name string     `json:"name"`
	}
	type Response struct {
		ID          uint32       `json:"id"`
		Code        types.Code   `json:"code"`
		Name        string       `json:"name"`
		Permissions []Permission `json:"permissions"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.RoleService.Update(ctx, role.UpdateReq{
		ID:              c.Param("id"),
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:   serviceResp.ID,
		Code: serviceResp.Code,
		Name: serviceResp.Name,
		Permissions: func() []Permission {
			perms := make([]Permission, len(serviceResp.Permissions))
			for index, perm := range serviceResp.Permissions {
				perms[index] = Permission{
					ID:   perm.ID,
					Code: perm.Code,
					Name: perm.Name,
				}
			}
			return perms
		}(),
	})
}

func (h HTTP) roleDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)
	//ctxID := contextid.Value(ctx)

	serviceErr := h.RoleService.Delete(ctx, role.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
