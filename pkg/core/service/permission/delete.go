package permission

import (
	"context"
	"errors"
	"strconv"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
)

type DeleteReq struct {
	ID string

	id uint32
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

	id, _ := strconv.ParseInt(req.ID, 10, 32)
	req.id = uint32(id)

	return nil
}

func (s Service) Delete(ctx context.Context, req DeleteReq) (serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return service.ErrInvalidInput(err.Error())
	}

	err = s.PermissionRepo.Delete(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return service.ErrDatasourceAccess("permission not found")
		default:
			return service.ErrDatasourceAccess("delete permission query error")
		}
	}

	return nil
}
