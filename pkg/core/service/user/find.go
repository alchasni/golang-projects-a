package user

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
	ID             uint64
	Username       string
	Email          string
	Password       string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
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

	user, err := s.UserRepo.Find(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("user not found")
		default:
			return resp, service.ErrDatasourceAccess("find user query error")
		}
	}

	return FindResp{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		Password:       user.Password,
		AvatarUrl:      user.AvatarUrl,
		OrganizationId: user.OrganizationId,
		FollowingCount: user.FollowingCount,
		FollowerCount:  user.FollowerCount,
	}, nil
}
