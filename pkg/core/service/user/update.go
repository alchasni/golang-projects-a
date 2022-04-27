package user

import (
	"context"
	"errors"
	"strconv"
	"time"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/useradapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type UpdateReq struct {
	ID       string
	Email    string
	RoleCode types.Code

	id uint64
}

type UpdateResp struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Role      *domain.Role
}

var (
	validatorTag_UpdateReqID    = validatoradapter.Tag().Required().Numeric()
	validatorTag_UpdateReqEmail = validatoradapter.Tag().Omitempty().Email()
)

func (req *UpdateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"id", req.ID, validatorTag_UpdateReqID},
		{"email", req.Email, validatorTag_UpdateReqEmail},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	id, _ := strconv.ParseInt(req.ID, 10, 64)
	req.id = uint64(id)

	return nil
}

func (s Service) Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	user, err := s.UserRepo.Update(ctx, req.id, useradapter.RepoUpdate{
		Email:    req.Email,
		RoleCode: req.RoleCode,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("user not found")
		default:
			return resp, service.ErrDatasourceAccess("update user query error")
		}
	}

	return UpdateResp{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}
