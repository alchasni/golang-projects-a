package role

import (
	"context"

	"twatter/pkg/core/adapter/roleadapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	"twatter/pkg/types"
)

type GetListReq struct {
	Code            types.Code
	Name            string
	PermissionCodes []types.Code
	PageNo          int
	PageSize        int
}

type GetListResp struct {
	Items    []domain.Role
	RowCount uint32
}

var (
	validatorTag_GetListReqPageNo   = validatoradapter.Tag().Omitempty().Gte(0)
	validatorTag_GetListReqPageSize = validatoradapter.Tag().Omitempty().Gte(0)
)

func (req *GetListReq) validate(v validatoradapter.Adapter) error {
	var err error

	fields := []validatoradapter.Field{
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

	roles, err := s.RoleRepo.GetList(ctx, roleadapter.RepoFilter{
		Code:            req.Code,
		Name:            req.Name,
		PermissionCodes: req.PermissionCodes,
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
	})
	if err != nil {
		return resp, service.ErrDatasourceAccess("get list role query error")
	}

	return GetListResp{
		Items:    roles.Items,
		RowCount: roles.RowCount,
	}, nil
}
