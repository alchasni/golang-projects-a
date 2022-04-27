package permission

import (
	"context"
	"errors"
	"strconv"

	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/permissionadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
	"golang-projects-a/pkg/types"
)

type UpdateReq struct {
	ID         string
	Name       string
	Restricted string

	id uint32
}

type UpdateResp struct {
	ID         uint32
	Code       types.Code
	Name       string
	Restricted string
}

var (
	validatorTag_UpdateReqID         = validatoradapter.Tag().Required().Numeric()
	validatorTag_UpdateReqRestricted = validatoradapter.Tag().Omitempty().OneOf(consts.Confirmation_Y, consts.Confirmation_N)
)

func (req *UpdateReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"id", req.ID, validatorTag_UpdateReqID},
		{"restricted", req.Restricted, validatorTag_UpdateReqRestricted},
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

func (s Service) Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	perm, err := s.PermissionRepo.Update(ctx, req.id, permissionadapter.RepoUpdate{
		Name:       req.Name,
		Restricted: req.Restricted,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrNotFound):
			return resp, service.ErrDatasourceAccess("permission not found")
		default:
			return resp, service.ErrDatasourceAccess("update permission query error")
		}
	}

	return UpdateResp{
		ID:         perm.ID,
		Code:       perm.Code,
		Name:       perm.Name,
		Restricted: perm.Restricted,
	}, nil
}
