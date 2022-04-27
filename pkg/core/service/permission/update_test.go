package permission

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/permissionadapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_permissionadapter "golang-projects-a/pkg/mocks/adapter/permissionadapter"
	"golang-projects-a/pkg/platform/validator"

	"github.com/golang/mock/gomock"
)

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepo := mock_permissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		PermissionRepo: permissionRepo,
		Validator:      v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := UpdateReq{
			ID:         "1",
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,

			id: 1,
		}
		repoUpdate := permissionadapter.RepoUpdate{
			Name:       req.Name,
			Restricted: req.Restricted,
		}
		perm := domain.Permission{
			ID:         req.id,
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       repoUpdate.Name,
			Restricted: repoUpdate.Restricted,
		}

		permissionRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(perm, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:         perm.ID,
			Code:       perm.Code,
			Name:       perm.Name,
			Restricted: perm.Restricted,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow no updating", func(t *testing.T) {
		req := UpdateReq{
			ID: "1",
			//Name:       "Permission Read",
			//Restricted: consts.Confirmation_Y,

			id: 1,
		}
		repoUpdate := permissionadapter.RepoUpdate{
			Name:       req.Name,
			Restricted: req.Restricted,
		}
		perm := domain.Permission{
			ID:         1,
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}

		permissionRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(perm, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:         perm.ID,
			Code:       perm.Code,
			Name:       perm.Name,
			Restricted: perm.Restricted,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := UpdateReq{
			//ID:         "1",
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,

			//id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := UpdateReq{
			ID:         "a",
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,

			id: 0,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return invalid input restricted oneof", func(t *testing.T) {
		req := UpdateReq{
			ID:         "1",
			Name:       "Permission Read",
			Restricted: "rand",

			id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on restricted, this field should be one of: Y N", e.Error())
	})

	t.Run("should return permission not found", func(t *testing.T) {
		req := UpdateReq{
			ID:         "1",
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,

			id: 1,
		}
		repoUpdate := permissionadapter.RepoUpdate{
			Name:       req.Name,
			Restricted: req.Restricted,
		}

		permissionRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Permission{}, adapter.ErrNotFound)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "permission not found", e.Error())
	})

	t.Run("should return update query error", func(t *testing.T) {
		req := UpdateReq{
			ID:         "1",
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,

			id: 1,
		}
		repoUpdate := permissionadapter.RepoUpdate{
			Name:       req.Name,
			Restricted: req.Restricted,
		}

		permissionRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Permission{}, adapter.ErrQuery)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "update permission query error", e.Error())
	})
}
