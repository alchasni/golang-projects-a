package rolepermission

import (
	"context"
	"errors"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type CheckExistenceReq struct {
	RoleCode       types.Code
	PermissionCode types.Code
}

type CheckExistenceResp struct{}

var (
	validatorTag_CheckExistenceReqRoleCode       = validatoradapter.Tag().Required()
	validatorTag_CheckExistenceReqPermissionCode = validatoradapter.Tag().Required()
)

func (req *CheckExistenceReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"role_code", req.RoleCode, validatorTag_CheckExistenceReqRoleCode},
		{"permission_code", req.PermissionCode, validatorTag_CheckExistenceReqPermissionCode},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) CheckExistence(ctx context.Context, req CheckExistenceReq) (serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	err = s.RolePermissionRepo.CheckExistence(ctx, req.RoleCode, req.PermissionCode)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return service.ErrDatasourceAccess("role permission doesn't exist")
		default:
			return service.ErrDatasourceAccess("check existence role permission query error")
		}
	}

	return nil
}
