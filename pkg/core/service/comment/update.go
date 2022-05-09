package comment

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"strconv"
)

type UpdateReq struct {
	ID             string
	Comment        string
	OrganizationId uint64

	id uint64
}

type UpdateResp struct {
	ID             uint64
	Comment        string
	OrganizationId uint64
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

	comment, err := s.CommentRepo.Update(ctx, req.id, commentadapter.RepoUpdate{
		Comment:        req.Comment,
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("comment not found")
		default:
			return resp, service.ErrDatasourceAccess("update comment query error")
		}
	}

	return UpdateResp{
		ID:             comment.ID,
		Comment:        comment.Comment,
		OrganizationId: comment.OrganizationId,
	}, nil
}
