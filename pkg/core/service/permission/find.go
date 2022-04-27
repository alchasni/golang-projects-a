package permission

import (
	"context"
	"errors"
	"strconv"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type FindReq struct {
	ID string

	id uint32
}

type FindResp struct {
	ID         uint32
	Code       types.Code
	Name       string
	Restricted string
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

	id, err := strconv.ParseInt(req.ID, 10, 32)
	req.id = uint32(id)

	return nil
}

func (s Service) Find(ctx context.Context, req FindReq) (resp FindResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	perm, err := s.PermissionRepo.Find(ctx, req.id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("permission not found")
		default:
			return resp, service.ErrDatasourceAccess("find permission query error")
		}
	}

	return FindResp{
		ID:         perm.ID,
		Code:       perm.Code,
		Name:       perm.Name,
		Restricted: perm.Restricted,
	}, nil
}
