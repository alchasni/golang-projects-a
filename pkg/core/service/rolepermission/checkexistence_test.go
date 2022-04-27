package rolepermission

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"twatter/pkg/consts"
	"twatter/pkg/core/adapter"
	"twatter/pkg/core/service"
	mock_rolepermissionadapter "twatter/pkg/mocks/adapter/rolepermissionadapter"
	"twatter/pkg/platform/validator"

	"github.com/golang/mock/gomock"
)

func TestService_CheckExistence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rolePermissionRepo := mock_rolepermissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RolePermissionRepo: rolePermissionRepo,
		Validator:          v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CheckExistenceReq{
			RoleCode:       consts.RoleCode_ROOT,
			PermissionCode: consts.PermissionCode_PermisisonRead,
		}

		rolePermissionRepo.EXPECT().CheckExistence(gomock.Any(), req.RoleCode, req.PermissionCode).Return(nil)

		e := svc.CheckExistence(context.Background(), req)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input role_code required", func(t *testing.T) {
		req := CheckExistenceReq{
			//RoleCode:       consts.RoleCode_ROOT,
			PermissionCode: consts.PermissionCode_PermisisonRead,
		}

		e := svc.CheckExistence(context.Background(), req)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on role_code, this field is required", e.Error())
	})

	t.Run("should return invalid input permission_code required", func(t *testing.T) {
		req := CheckExistenceReq{
			RoleCode: consts.RoleCode_ROOT,
			//PermissionCode: consts.PermissionCode_PermisisonRead,
		}

		e := svc.CheckExistence(context.Background(), req)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on permission_code, this field is required", e.Error())
	})

	t.Run("should return role permission doesn't exist", func(t *testing.T) {
		req := CheckExistenceReq{
			RoleCode:       consts.RoleCode_ROOT,
			PermissionCode: consts.PermissionCode_PermisisonRead,
		}

		rolePermissionRepo.EXPECT().CheckExistence(gomock.Any(), req.RoleCode, req.PermissionCode).Return(adapter.ErrNotFound)

		e := svc.CheckExistence(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "role permission doesn't exist", e.Error())
	})

	t.Run("should return check existence query error", func(t *testing.T) {
		req := CheckExistenceReq{
			RoleCode:       consts.RoleCode_ROOT,
			PermissionCode: consts.PermissionCode_PermisisonRead,
		}

		rolePermissionRepo.EXPECT().CheckExistence(gomock.Any(), req.RoleCode, req.PermissionCode).Return(adapter.ErrQuery)

		e := svc.CheckExistence(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "check existence role permission query error", e.Error())
	})
}
