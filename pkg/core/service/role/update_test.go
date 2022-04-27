package role

import (
	"context"
	"testing"

	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/roleadapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_roleadapter "golang-projects-a/pkg/mocks/adapter/roleadapter"
	"golang-projects-a/pkg/platform/validator"
	"golang-projects-a/pkg/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roleRepo := mock_roleadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RoleRepo:  roleRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := UpdateReq{
			ID:              "1",
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			id: 1,
		}
		repoUpdate := roleadapter.RepoUpdate{
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
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

		roleRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(role, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Permissions: role.Permissions,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow no updating", func(t *testing.T) {
		req := UpdateReq{
			ID: "1",
			//Name:            "Root",
			//PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			id: 1,
		}
		repoUpdate := roleadapter.RepoUpdate{
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
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

		roleRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(role, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Permissions: role.Permissions,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := UpdateReq{
			//ID:              "1",
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			//id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := UpdateReq{
			ID:              "a",
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			id: 0,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return role not found", func(t *testing.T) {
		req := UpdateReq{
			ID:              "1",
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			id: 1,
		}
		repoUpdate := roleadapter.RepoUpdate{
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}

		roleRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Role{}, adapter.ErrNotFound)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "role not found", e.Error())
	})

	t.Run("should return update query error", func(t *testing.T) {
		req := UpdateReq{
			ID:              "1",
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},

			id: 1,
		}
		repoUpdate := roleadapter.RepoUpdate{
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}

		roleRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Role{}, adapter.ErrQuery)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "update role query error", e.Error())
	})
}
