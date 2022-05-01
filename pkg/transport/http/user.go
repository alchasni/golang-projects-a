package http

import (
	"golang-projects-a/pkg/core/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/transport/http/middleware"
)

func (h HTTP) userFind(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Response struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	serviceResp, serviceErr := h.UserService.Find(ctx, user.FindReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:       serviceResp.ID,
		Username: serviceResp.Username,
		Email:    serviceResp.Email,
		Password: serviceResp.Password,
	})
}

func (h HTTP) userGetList(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		ID        uint64 `query:"id"`
		Username  string `query:"username"`
		Email     string `query:"email"`
		Password  string `query:"password"`
		AvatarUrl string `query:"avatar_url"`

		Limit  int `query:"limit"`
		Offset int `query:"offset"`
	}
	type User struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	}
	type Response struct {
		Items    []User `json:"items"`
		RowCount uint64 `json:"row_count"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.UserService.GetList(ctx, user.GetListReq{
		ID:        req.ID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		Items: func() []User {
			items := make([]User, len(serviceResp.Items))
			for index, item := range serviceResp.Items {
				items[index] = User{
					ID:       item.ID,
					Username: item.Username,
				}
			}
			return items
		}(),
		RowCount: serviceResp.RowCount,
	})
}

func (h HTTP) userCreate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type Response struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.UserService.Create(ctx, user.CreateReq{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: "asdasd",
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:       serviceResp.ID,
		Username: serviceResp.Username,
	})
}

//func (h HTTP) permissionUpdate(c echo.Context) error {
//	ctx := middleware.ContextID(c)
//	//ctxID := contextid.Value(ctx)
//
//	type Request struct {
//		Name       string `json:"name"`
//		Restricted string `json:"restricted"`
//	}
//	type Response struct {
//		ID         uint32     `json:"id"`
//		Code       types.Code `json:"code"`
//		Name       string     `json:"name"`
//		Restricted string     `json:"restricted"`
//	}
//
//	req := Request{}
//	if err := c.Bind(&req); err != nil {
//		return service.ErrInvalidInput(err.Error())
//	}
//
//	serviceResp, serviceErr := h.PermissionService.Update(ctx, permission.UpdateReq{
//		ID:         c.Param("id"),
//		Name:       req.Name,
//		Restricted: req.Restricted,
//	})
//	if serviceErr != nil {
//		return serviceErr
//	}
//
//	return c.JSON(http.StatusOK, Response{
//		ID:         serviceResp.ID,
//		Code:       serviceResp.Code,
//		Name:       serviceResp.Name,
//		Restricted: serviceResp.Restricted,
//	})
//}

func (h HTTP) userDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)

	serviceErr := h.UserService.Delete(ctx, user.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
