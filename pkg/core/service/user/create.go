package user

import (
	"context"
	"errors"
	"time"

	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/types"
)

type CreateReq struct {
	Username string
	Email    string
	Password string
	RoleCode types.Code
}

type CreateResp struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Role      *domain.Role
}

var (
	validatorTag_CreateReqUsername = validatoradapter.Tag().Required().AlphaNum()
	validatorTag_CreateReqEmail    = validatoradapter.Tag().Omitempty().Email()
	validatorTag_CreateReqPassword = validatoradapter.Tag().Required()
	validatorTag_CreateReqRoleCode = validatoradapter.Tag().Required()
)

func (req CreateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"username", req.Username, validatorTag_CreateReqUsername},
		{"email", req.Email, validatorTag_CreateReqEmail},
		{"password", req.Password, validatorTag_CreateReqPassword},
		{"role_code", req.RoleCode, validatorTag_CreateReqRoleCode},
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

	user, err := s.UserRepo.Create(ctx, useradapter.RepoCreate{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		RoleCode: req.RoleCode,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrDuplicate):
			return resp, service.ErrDatasourceAccess("username is already exist")
		default:
			return resp, service.ErrDatasourceAccess("create user query error")
		}
	}

	return CreateResp{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}
