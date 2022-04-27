package role

import (
	"context"
	"errors"
	"strconv"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/roleadapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type UpdateReq struct {
	ID              string
	Name            string
	PermissionCodes []types.Code

	id uint32
}

type UpdateResp struct {
	ID          uint32
	Code        types.Code
	Name        string
	Permissions []domain.Permission
}

var (
	validatorTag_UpdateReqID = validatoradapter.Tag().Required().Numeric()
)

func (req *UpdateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"id", req.ID, validatorTag_UpdateReqID},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	id, _ := strconv.ParseInt(req.ID, 10, 32)
	req.id = uint32(id)

	return nil
}

func (s Service) Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	role, err := s.RoleRepo.Update(ctx, req.id, roleadapter.RepoUpdate{
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("role not found")
		default:
			return resp, service.ErrDatasourceAccess("update role query error")
		}
	}

	return UpdateResp{
		ID:          role.ID,
		Code:        role.Code,
		Name:        role.Name,
		Permissions: role.Permissions,
	}, nil
}
