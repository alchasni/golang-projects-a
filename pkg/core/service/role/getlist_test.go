package role

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"twatter/pkg/consts"
	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/roleadapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	mock_roleadapter "twatter/pkg/mocks/adapter/roleadapter"
	"twatter/pkg/platform/validator"
	"twatter/pkg/types"

	"github.com/golang/mock/gomock"
)

func TestService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roleRepo := mock_roleadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RoleRepo:  roleRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := GetListReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoFilter := roleadapter.RepoFilter{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}
		roles := domain.Roles{
			Items: []domain.Role{
				{
					ID:   1,
					Code: consts.RoleCode_ROOT,
					Name: "Root",
					Permissions: []domain.Permission{
						{
							ID:         1,
							Code:       consts.PermissionCode_PermisisonRead,
							Name:       "Permission Read",
							Restricted: consts.Confirmation_Y,
						},
					},
				},
			},
			RowCount: 1,
		}

		roleRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(roles, nil)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{
			Items:    roles.Items,
			RowCount: roles.RowCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input page_no greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
			PageNo:          -1,
			PageSize:        50,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on page_no, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return invalid input page_size greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
			PageNo:          1,
			PageSize:        -50,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on page_size, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return get list query error", func(t *testing.T) {
		req := GetListReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoFilter := roleadapter.RepoFilter{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}

		roleRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(domain.Roles{}, adapter.ErrQuery)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "get list role query error", e.Error())
	})
}
