package organization

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"strconv"
)

type UpdateReq struct {
	ID   string
	Name string

	id uint64
}

type UpdateResp struct {
	ID    uint64
	Name  string
	Email string
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

	id, _ := strconv.ParseUint(req.ID, 10, 64)
	req.id = id

	return nil
}

func (s Service) Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	user, err := s.OrganizationRepo.Update(ctx, req.id, organizationadapter.RepoUpdate{
		Name: req.Name,
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
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
