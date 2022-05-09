package user

import (
	"context"
	"crypto/md5"
	"fmt"
	"golang-projects-a/pkg/core/adapter/useradapter"

	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type CreateReq struct {
	Username       string
	Email          string
	Password       string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
}

type CreateResp struct {
	ID             uint64
	Username       string
	Email          string
	AvatarUrl      string
	OrganizationId uint64
	FollowingCount uint64
	FollowerCount  uint64
}

var (
	validatorTag_CreateReqUsername       = validatoradapter.Tag().Required()
	validatorTag_CreateReqEmail          = validatoradapter.Tag().Required().Email()
	validatorTag_CreateReqPassword       = validatoradapter.Tag().Required()
	validatorTag_CreateReqAvatarUrl      = validatoradapter.Tag().Required()
	validatorTag_CreateReqOrganizationId = validatoradapter.Tag().Required().Numeric()
	validatorTag_CreateReqFollowingCount = validatoradapter.Tag().Required().Numeric()
	validatorTag_CreateReqFollowerCount  = validatoradapter.Tag().Required().Numeric()
)

func (req *CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"username", req.Username, validatorTag_CreateReqUsername},
		{"email", req.Email, validatorTag_CreateReqEmail},
		{"password", req.Password, validatorTag_CreateReqPassword},
		{"avatar_url", req.AvatarUrl, validatorTag_CreateReqAvatarUrl},
		{"organization_id", req.OrganizationId, validatorTag_CreateReqOrganizationId},
		{"following_count", req.FollowingCount, validatorTag_CreateReqFollowingCount},
		{"follower_count", req.FollowerCount, validatorTag_CreateReqFollowerCount},
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

	data := []byte(req.Password)
	password := fmt.Sprintf("%x", md5.Sum(data))

	user, err := s.UserRepo.Create(ctx, useradapter.RepoCreate{
		Username:       req.Username,
		Email:          req.Email,
		Password:       password,
		AvatarUrl:      req.AvatarUrl,
		OrganizationId: req.OrganizationId,
		FollowingCount: req.FollowingCount,
		FollowerCount:  req.FollowerCount,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("create user query error")
	}

	return CreateResp{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		AvatarUrl:      user.AvatarUrl,
		OrganizationId: user.OrganizationId,
		FollowingCount: user.FollowingCount,
		FollowerCount:  user.FollowerCount,
	}, nil
}
