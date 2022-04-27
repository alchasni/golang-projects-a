package permission

import (
	"context"
	"testing"

	"twatter/pkg/consts"
	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/permissionadapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	mock_permissionadapter "twatter/pkg/mocks/adapter/permissionadapter"
	"twatter/pkg/platform/validator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepo := mock_permissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		PermissionRepo: permissionRepo,
		Validator:      v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := GetListReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: "Y",
			PageNo:     1,
			PageSize:   50,
		}
		repoFilter := permissionadapter.RepoFilter{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
		}
		perms := domain.Permissions{
			Items: []domain.Permission{
				{
					ID:         1,
					Code:       consts.PermissionCode_PermisisonRead,
					Name:       "Permission Read",
					Restricted: consts.Confirmation_Y,
				},
			},
			RowCount: 1,
		}

		permissionRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(perms, nil)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{
			Items:    perms.Items,
			RowCount: perms.RowCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input restricted oneof", func(t *testing.T) {
		req := GetListReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: "rand",
			PageNo:     1,
			PageSize:   50,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on restricted, this field should be one of: Y N", e.Error())
	})

	t.Run("should return invalid input page_no greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
			PageNo:     -1,
			PageSize:   50,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on page_no, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return invalid input page_size greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
			PageNo:     1,
			PageSize:   -50,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on page_size, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return get list query error", func(t *testing.T) {
		req := GetListReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
			PageNo:     1,
			PageSize:   50,
		}
		repoFilter := permissionadapter.RepoFilter{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
			PageNo:     req.PageNo,
			PageSize:   req.PageSize,
		}

		permissionRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(domain.Permissions{}, adapter.ErrQuery)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "get list permission query error", e.Error())
	})
}
