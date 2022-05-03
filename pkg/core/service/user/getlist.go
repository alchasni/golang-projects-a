package user

import (
	"context"
	"golang-projects-a/pkg/core/adapter/useradapter"

	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
)

type GetListReq struct {
	ID             uint64
	Username       string
	Email          string
	Password       string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64

	Limit  int64
	Offset int64
}

type GetListResp struct {
	Items    []domain.User
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

	user, err := s.UserRepo.GetList(ctx, useradapter.RepoFilter{
		ID:             req.ID,
		Username:       req.Username,
		Email:          req.Email,
		Password:       req.Password,
		AvatarUrl:      req.AvatarUrl,
		OrganizationId: req.OrganizationId,
		FollowingCount: req.FollowingCount,
		FollowerCount:  req.FollowerCount,
		Limit:          req.Limit,
		Offset:         req.Offset,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list user query error")
	}

	return GetListResp{
		Items:    user.Items,
		RowCount: user.RowCount,
	}, nil
}
