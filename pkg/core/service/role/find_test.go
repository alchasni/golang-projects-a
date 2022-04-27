package role

import (
	"context"
	"testing"

	"twatter/pkg/consts"
	"twatter/pkg/core/adapter"
	"twatter/pkg/core/domain"
	"twatter/pkg/core/service"
	mock_roleadapter "twatter/pkg/mocks/adapter/roleadapter"
	"twatter/pkg/platform/validator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roleRepo := mock_roleadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RoleRepo:  roleRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}
		role := domain.Role{
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
		}

		roleRepo.EXPECT().Find(gomock.Any(), req.id).Return(role, nil)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{
			ID:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Permissions: role.Permissions,
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

	t.Run("should return role not found", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		roleRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Role{}, adapter.ErrNotFound)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "role not found", e.Error())
	})

	t.Run("should return find query error", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		roleRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Role{}, adapter.ErrQuery)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "find role query error", e.Error())
	})
}
