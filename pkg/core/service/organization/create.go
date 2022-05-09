package organization

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type CreateReq struct {
	Name string
}

type CreateResp struct {
	ID   uint64
	Name string
}

var (
	validatorTag_CreateReqName = validatoradapter.Tag().AlphaNum().Required()
)

func (req *CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
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

	organization, err := s.OrganizationRepo.Create(ctx, organizationadapter.RepoCreate{
		Name: req.Name,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrDuplicate):
			return resp, service.ErrDatasourceAccess("duplicate organization name")
		default:
			return resp, service.ErrDatasourceAccess("create organization query error")
		}
	}

	return CreateResp{
		ID:   organization.ID,
		Name: organization.Name,
	}, nil
}
