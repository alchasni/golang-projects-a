package user

import (
	"context"
	"errors"
	"strconv"
	"time"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
)

type FindReq struct {
	ID string

	id uint64
}

type FindResp struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Role      *domain.Role
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

	id, err := strconv.ParseInt(req.ID, 10, 64)
	req.id = uint64(id)

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
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}
