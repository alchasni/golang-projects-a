package permission

import (
	"context"

	"twatter/pkg/consts"
	"twatter/pkg/core/adapter/permissionadapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type GetListReq struct {
	Code       types.Code
	Name       string
	Restricted string
	PageNo     int
	PageSize   int
}

type GetListResp struct {
	Items    []domain.Permission
	RowCount uint32
}

var (
	validatorTag_GetListReqRestricted = validatoradapter.Tag().Omitempty().OneOf(consts.Confirmation_Y, consts.Confirmation_N)
	validatorTag_GetListReqPageNo     = validatoradapter.Tag().Omitempty().Gte(0)
	validatorTag_GetListReqPageSize   = validatoradapter.Tag().Omitempty().Gte(0)
)

func (req *GetListReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
		{"restricted", req.Restricted, validatorTag_GetListReqRestricted},
		{"page_no", req.PageNo, validatorTag_GetListReqPageNo},
		{"page_size", req.PageSize, validatorTag_GetListReqPageSize},
	}
	for _, field := range fields {
		if err = v.Var(field); err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetList(ctx context.Context, req GetListReq) (resp GetListResp, serviceErr service.Error) {
	err := req.validate(s.Validator)
	if err != nil {
		return resp, service.ErrInvalidInput(err.Error())
	}

	perms, err := s.PermissionRepo.GetList(ctx, permissionadapter.RepoFilter{
		Code:       req.Code,
		Name:       req.Name,
		Restricted: req.Restricted,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list permission query error")
	}

	return GetListResp{
		Items:    perms.Items,
		RowCount: perms.RowCount,
	}, nil
}
