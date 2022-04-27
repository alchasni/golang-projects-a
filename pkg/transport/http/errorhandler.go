package http

import (
	"fmt"
	"net/http"
	"strings"

	"twatter/pkg/core/service"
	internal_validator "twatter/pkg/platform/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrHandler struct {
	E *echo.Echo
}

type resp struct {
	ErrorObj errorObj `json:"error"`
}

type errorObj struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func (ce ErrHandler) Handler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
	)

	errObj := errorObj{}
	switch e := err.(type) {
	case service.Error:
		errObj.Code = e.Code()
		errObj.Message = e.Error()
		code = GetServiceErrorStatusCode(e)
	case validator.ValidationErrors:
		var errMsg []string
		for _, v := range e {
			errMsg = append(errMsg, fmt.Sprintf("invalid value on %s, %s", v.Field(), internal_validator.ErrorReason(v)))
		}
		errObj.Code = service.ErrorCode_InvalidInput
		errObj.Message = strings.Join(errMsg, ",")
		code = http.StatusBadRequest
	default:
		errObj.Code = service.ErrorCode_General
		errObj.Message = e.Error()
	}

	r := resp{ErrorObj: errObj}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, r)
		}
		if err != nil {
			ce.E.Logger.Error(err)
		}
	}
}

func GetServiceErrorStatusCode(serviceErr service.Error) int {
	switch serviceErr.Code() {
	case service.ErrorCode_InvalidInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
