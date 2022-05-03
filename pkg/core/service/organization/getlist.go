package organization

import (
	"context"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
)

type GetListReq struct {
	ID   uint64
	Name string

	Limit  int
	Offset int
}

type GetListResp struct {
	Items    []domain.Organization
	RowCount uint64
}

var (
	validatorTag_GetListLimit  = validatoradapter.Tag().Omitempty().Gte(0)
	validatorTag_GetListOffset = validatoradapter.Tag().Omitempty().Gte(0)
)

func (req *GetListReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"limit", req.Limit, validatorTag_GetListLimit},
		{"offset", req.Offset, validatorTag_GetListOffset},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetList(ctx context.Context, req GetListReq) (resp GetListResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	organization, err := s.OrganizationRepo.GetList(ctx, organizationadapter.RepoFilter{
		ID:     req.ID,
		Name:   req.Name,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list organization query error")
	}

	return GetListResp{
		Items:    organization.Items,
		RowCount: organization.RowCount,
	}, nil
}
