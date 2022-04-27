package permission

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/service"
	mock_permissionadapter "golang-projects-a/pkg/mocks/adapter/permissionadapter"
	"golang-projects-a/pkg/platform/validator"

	"github.com/golang/mock/gomock"
)

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepo := mock_permissionadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		PermissionRepo: permissionRepo,
		Validator:      v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		permissionRepo.EXPECT().Delete(gomock.Any(), req.id).Return(nil)

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := DeleteReq{
			//ID: 1,

			//id: 1,
		}

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := DeleteReq{
			ID: "a",

			id: 0,
		}

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return permission not found", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		permissionRepo.EXPECT().Delete(gomock.Any(), req.id).Return(adapter.ErrNotFound)

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "permission not found", e.Error())
	})

	t.Run("should return delete query error", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		permissionRepo.EXPECT().Delete(gomock.Any(), req.id).Return(adapter.ErrQuery)

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "delete permission query error", e.Error())
	})
}
