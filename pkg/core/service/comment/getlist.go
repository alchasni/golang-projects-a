package comment

import (
	"context"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
)

type GetListReq struct {
	ID             uint64
	OrganizationId uint64

	Limit  int
	Offset int
}

type GetListResp struct {
	Items    []domain.Comment
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

	comment, err := s.CommentRepo.GetList(ctx, commentadapter.RepoFilter{
		ID:             req.ID,
		OrganizationId: req.OrganizationId,
		Limit:          req.Limit,
		Offset:         req.Offset,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list comment query error")
	}

	return GetListResp{
		Items:    comment.Items,
		RowCount: comment.RowCount,
	}, nil
}
