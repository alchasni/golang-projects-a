package comment

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type CreateReq struct {
	Comment        string
	OrganizationId uint64
}

type CreateResp struct {
	ID             uint64
	Comment        string
	OrganizationId uint64
}

var (
	validatorTag_CreateReqComment        = validatoradapter.Tag().Required()
	validatorTag_CreateReqOrganizationId = validatoradapter.Tag().Required().Numeric()
)

func (req *CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"comment", req.Comment, validatorTag_CreateReqComment},
		{"organization_id", req.OrganizationId, validatorTag_CreateReqOrganizationId},
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

	comment, err := s.CommentRepo.Create(ctx, commentadapter.RepoCreate{
		Comment:        req.Comment,
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrDuplicate):
			return resp, service.ErrDatasourceAccess("duplicate comment code")
		default:
			return resp, service.ErrDatasourceAccess("create comment query error")
		}
	}

	return CreateResp{
		ID:             comment.ID,
		Comment:        comment.Comment,
		OrganizationId: comment.OrganizationId,
	}, nil
}
