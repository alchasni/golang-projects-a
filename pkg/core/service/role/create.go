package role

import (
	"context"
	"errors"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/roleadapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type CreateReq struct {
	Code            types.Code
	Name            string
	PermissionCodes []types.Code
}

type CreateResp struct {
	ID          uint32
	Code        types.Code
	Name        string
	Permissions []domain.Permission
}

var (
	validatorTag_CreateReqCode = validatoradapter.Tag().Required()
	validatorTag_CreateReqName = validatoradapter.Tag().Required()
)

func (req *CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"code", req.Code, validatorTag_CreateReqCode},
		{"name", req.Name, validatorTag_CreateReqName},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) Create(ctx context.Context, req CreateReq) (resp CreateResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	role, err := s.RoleRepo.Create(ctx, roleadapter.RepoCreate{
		Code:            req.Code,
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrDuplicate):
			return resp, service.ErrDatasourceAccess("duplicate role code")
		default:
			return resp, service.ErrDatasourceAccess("create role query error")
		}
	}

	return CreateResp{
		ID:          role.ID,
		Code:        role.Code,
		Name:        role.Name,
		Permissions: role.Permissions,
	}, nil
}
