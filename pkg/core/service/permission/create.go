package permission

import (
	"context"
	"errors"

	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/permissionadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/types"
)

type CreateReq struct {
	Code       types.Code
	Name       string
	Restricted string
}

type CreateResp struct {
	ID         uint32
	Code       types.Code
	Name       string
	Restricted string
}

var (
	validatorTag_CreateReqCode       = validatoradapter.Tag().Required()
	validatorTag_CreateReqName       = validatoradapter.Tag().Required()
	validatorTag_CreateReqRestricted = validatoradapter.Tag().Omitempty().OneOf(consts.Confirmation_Y, consts.Confirmation_N)
)

func (req *CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"code", req.Code, validatorTag_CreateReqCode},
		{"name", req.Name, validatorTag_CreateReqName},
		{"restricted", req.Restricted, validatorTag_CreateReqRestricted},
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

	perm, err := s.PermissionRepo.Create(ctx, permissionadapter.RepoCreate{
		Code:       req.Code,
		Name:       req.Name,
		Restricted: req.Restricted,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrDuplicate):
			return resp, service.ErrDatasourceAccess("duplicate permission code")
		default:
			return resp, service.ErrDatasourceAccess("create permission query error")
		}
	}

	return CreateResp{
		ID:         perm.ID,
		Code:       perm.Code,
		Name:       perm.Name,
		Restricted: perm.Restricted,
	}, nil
}
