package user

import (
	"context"

	"twatter/pkg/core/adapter/useradapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type GetListReq struct {
	Username string
	Email    string
	RoleCode types.Code
	PageNo   int
	PageSize int
}

type GetListResp struct {
	Items    []domain.User
	RowCount uint64
}

var (
	validatorTag_GetListReqPageNo   = validatoradapter.Tag().Omitempty().Gte(0)
	validatorTag_GetListReqPageSize = validatoradapter.Tag().Omitempty().Gte(0)
)

func (req GetListReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"page_no", req.PageNo, validatorTag_GetListReqPageNo},
		{"page_size", req.PageSize, validatorTag_GetListReqPageSize},
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

	users, err := s.UserRepo.GetList(ctx, useradapter.RepoFilter{
		Username: req.Username,
		Email:    req.Email,
		RoleCode: req.RoleCode,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list user query error")
	}

	return GetListResp{
		Items:    users.Items,
		RowCount: users.RowCount,
	}, nil
}
