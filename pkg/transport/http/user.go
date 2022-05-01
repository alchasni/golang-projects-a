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
		ID        uint64 `json:"id"`
		Username  string `json:"username"`
		Email     string `query:"email"`
		Password  string `query:"password"`
		AvatarUrl string `query:"avatar_url"`
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
					ID:        item.ID,
					Username:  item.Username,
					Email:     item.Email,
					Password:  item.Password,
					AvatarUrl: item.AvatarUrl,
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
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		AvatarUrl string `json:"avatar_url"`
	}
	type Response struct {
		ID        uint64 `json:"id"`
		Username  string `json:"username"`
		AvatarUrl string `json:"avatar_url"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.UserService.Create(ctx, user.CreateReq{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:        serviceResp.ID,
		Username:  serviceResp.Username,
		AvatarUrl: serviceResp.AvatarUrl,
	})
}

func (h HTTP) userUpdate(c echo.Context) error {
	ctx := middleware.ContextID(c)

	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	type Response struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	req := Request{}
	if err := c.Bind(&req); err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	serviceResp, serviceErr := h.UserService.Update(ctx, user.UpdateReq{
		ID:       c.Param("id"),
		Username: req.Username,
		Email:    req.Email,
	})
	if serviceErr != nil {
		return serviceErr
	}

	return c.JSON(http.StatusOK, Response{
		ID:       serviceResp.ID,
		Username: serviceResp.Username,
		Email:    serviceResp.Email,
	})
}

func (h HTTP) userDelete(c echo.Context) error {
	ctx := middleware.ContextID(c)

	serviceErr := h.UserService.Delete(ctx, user.DeleteReq{ID: c.Param("id")})
	if serviceErr != nil {
		return serviceErr
	}

	return c.NoContent(http.StatusOK)
}
