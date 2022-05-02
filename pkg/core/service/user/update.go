package user

import (
	"context"
	"errors"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"strconv"
)

type UpdateReq struct {
	ID             string
	Email          string
	Username       string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64

	id uint64
}

type UpdateResp struct {
	ID             uint64
	Username       string
	Email          string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
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

	user, err := s.UserRepo.Update(ctx, req.id, useradapter.RepoUpdate{
		Username:       req.Username,
		Email:          req.Email,
		OrganizationId: req.OrganizationId,
		FollowingCount: req.FollowingCount,
		FollowerCount:  req.FollowerCount,
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
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		OrganizationId: user.OrganizationId,
		FollowingCount: user.FollowingCount,
		FollowerCount:  user.FollowerCount,
	}, nil
}
