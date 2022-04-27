package role

import (
	"context"
	"errors"
	"strconv"

	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/types"
)

type FindReq struct {
	ID string

	id uint32
}

type FindResp struct {
	ID          uint32
	Code        types.Code
	Name        string
	Permissions []domain.Permission
}

var (
	validatorTag_FindReqID = validatoradapter.Tag().Required().Numeric()
)

func (req *FindReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"id", req.ID, validatorTag_FindReqID},
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

func (s Service) Find(ctx context.Context, req FindReq) (resp FindResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	role, err := s.RoleRepo.Find(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("role not found")
		default:
			return resp, service.ErrDatasourceAccess("find role query error")
		}
	}

	return FindResp{
		ID:          role.ID,
		Code:        role.Code,
		Name:        role.Name,
		Permissions: role.Permissions,
	}, nil
}
