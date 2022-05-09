package organization

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type FindByNameReq struct {
	Name string
}

type FindByNameResp struct {
	ID uint64
}

var (
	validatorTag_FindByNameReqName = validatoradapter.Tag().Required()
)

func (req *FindByNameReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"name", req.Name, validatorTag_FindByNameReqName},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) FindByName(ctx context.Context, req FindByNameReq) (resp FindByNameResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	organization, err := s.OrganizationRepo.FindByName(ctx, req.Name)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("organization not found")
		default:
			return resp, service.ErrDatasourceAccess("find organization query error")
		}
	}

	return FindByNameResp{
		ID: organization.ID,
	}, nil
}
