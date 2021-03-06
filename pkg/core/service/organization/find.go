package organization

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"strconv"
)

type FindReq struct {
	ID string

	id uint64
}

type FindResp struct {
	ID   uint64
	Name string
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

	id, _ := strconv.ParseUint(req.ID, 10, 64)
	req.id = id

	return nil
}

func (s Service) Find(ctx context.Context, req FindReq) (resp FindResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	organization, err := s.OrganizationRepo.Find(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("organization not found")
		default:
			return resp, service.ErrDatasourceAccess("find organization query error")
		}
	}

	return FindResp{
		ID:   organization.ID,
		Name: organization.Name,
	}, nil
}
