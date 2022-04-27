package user

import (
	"context"
	"errors"
	"strconv"

	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type DeleteReq struct {
	ID string

	id uint64
}

type DeleteResp struct{}

var (
	validatorTag_DeleteReqID = validatoradapter.Tag().Required().Numeric()
)

func (req *DeleteReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"id", req.ID, validatorTag_DeleteReqID},
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

func (s Service) Delete(ctx context.Context, req DeleteReq) (serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	err = s.UserRepo.Delete(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return service.ErrDatasourceAccess("user not found")
		default:
			return service.ErrDatasourceAccess("delete user query error")
		}
	}

	return nil
}
