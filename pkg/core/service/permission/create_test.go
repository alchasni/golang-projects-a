package permission

import (
	"context"
	"testing"

	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/permissionadapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_permissionadapter "golang-projects-a/pkg/mocks/adapter/permissionadapter"
	"golang-projects-a/pkg/platform/validator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepo := mock_permissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		PermissionRepo: permissionRepo,
		Validator:      v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CreateReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}
		repoCreate := permissionadapter.RepoCreate{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
		}
		perm := domain.Permission{
			ID:         1,
			Code:       repoCreate.Code,
			Name:       repoCreate.Name,
			Restricted: repoCreate.Restricted,
		}

		permissionRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(perm, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:         perm.ID,
			Code:       perm.Code,
			Name:       perm.Name,
			Restricted: perm.Restricted,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow omitempty", func(t *testing.T) {
		req := CreateReq{
			Code: consts.PermissionCode_PermisisonRead,
			Name: "Permission Read",
			//Restricted: consts.Confirmation_Y,
		}
		repoCreate := permissionadapter.RepoCreate{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
		}
		perm := domain.Permission{
			ID:         1,
			Code:       repoCreate.Code,
			Name:       repoCreate.Name,
			Restricted: consts.Confirmation_Y,
		}

		permissionRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(perm, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:         perm.ID,
			Code:       perm.Code,
			Name:       perm.Name,
			Restricted: perm.Restricted,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input code required", func(t *testing.T) {
		req := CreateReq{
			//Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on code, this field is required", e.Error())
	})

	t.Run("should return invalid input name required", func(t *testing.T) {
		req := CreateReq{
			Code: consts.PermissionCode_PermisisonRead,
			//Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on name, this field is required", e.Error())
	})

	t.Run("should return invalid input restricted oneof", func(t *testing.T) {
		req := CreateReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: "rand",
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on restricted, this field should be one of: Y N", e.Error())
	})

	t.Run("should return duplicate code", func(t *testing.T) {
		req := CreateReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}
		repoCreate := permissionadapter.RepoCreate{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
		}

		permissionRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Permission{}, adapter.ErrDuplicate)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "duplicate permission code", e.Error())
	})

	t.Run("should return create query error", func(t *testing.T) {
		req := CreateReq{
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}
		repoCreate := permissionadapter.RepoCreate{
			Code:       req.Code,
			Name:       req.Name,
			Restricted: req.Restricted,
		}

		permissionRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Permission{}, adapter.ErrQuery)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "create permission query error", e.Error())
	})
}
