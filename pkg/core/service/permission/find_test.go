package permission

import (
	"context"
	"testing"

	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_permissionadapter "golang-projects-a/pkg/mocks/adapter/permissionadapter"
	"golang-projects-a/pkg/platform/validator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepo := mock_permissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		PermissionRepo: permissionRepo,
		Validator:      v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}
		perm := domain.Permission{
			ID:         req.id,
			Code:       consts.PermissionCode_PermisisonRead,
			Name:       "Permission Read",
			Restricted: consts.Confirmation_Y,
		}

		permissionRepo.EXPECT().Find(gomock.Any(), req.id).Return(perm, nil)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{
			ID:         perm.ID,
			Code:       perm.Code,
			Name:       perm.Name,
			Restricted: perm.Restricted,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := FindReq{
			//ID: "1",

			//id: 1,
		}

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := FindReq{
			ID: "a",

			id: 0,
		}

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return permission not found", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		permissionRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Permission{}, adapter.ErrNotFound)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "permission not found", e.Error())
	})

	t.Run("should return find query error", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		permissionRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Permission{}, adapter.ErrQuery)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "find permission query error", e.Error())
	})
}
