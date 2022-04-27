package role

import (
	"context"
	"testing"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/service"
	mock_roleadapter "twatter/pkg/mocks/adapter/roleadapter"
	"twatter/pkg/platform/validator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roleRepo := mock_roleadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RoleRepo:  roleRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		roleRepo.EXPECT().Delete(gomock.Any(), req.id).Return(nil)

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

	t.Run("should return role not found", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		roleRepo.EXPECT().Delete(gomock.Any(), req.id).Return(adapter.ErrNotFound)

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "role not found", e.Error())
	})

	t.Run("should return delete query error", func(t *testing.T) {
		req := DeleteReq{
			ID: "1",

			id: 1,
		}

		roleRepo.EXPECT().Delete(gomock.Any(), req.id).Return(adapter.ErrQuery)

		e := svc.Delete(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "delete role query error", e.Error())
	})
}
