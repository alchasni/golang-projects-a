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

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roleRepo := mock_roleadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		RoleRepo:  roleRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CreateReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoCreate := roleadapter.RepoCreate{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}
		role := domain.Role{
			ID:   1,
			Code: repoCreate.Code,
			Name: repoCreate.Name,
			Permissions: []domain.Permission{
				{
					ID:         1,
					Code:       consts.PermissionCode_PermisisonRead,
					Name:       "Permission Read",
					Restricted: consts.Confirmation_Y,
				},
			},
		}

		roleRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(role, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Permissions: role.Permissions,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow omitempty", func(t *testing.T) {
		req := CreateReq{
			Code: consts.RoleCode_ROOT,
			Name: "Root",
			//PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoCreate := roleadapter.RepoCreate{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}
		role := domain.Role{
			ID:          1,
			Code:        repoCreate.Code,
			Name:        repoCreate.Name,
			Permissions: []domain.Permission{},
		}

		roleRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(role, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Permissions: role.Permissions,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input code required", func(t *testing.T) {
		req := CreateReq{
			//Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on code, this field is required", e.Error())
	})

	t.Run("should return invalid input code required", func(t *testing.T) {
		req := CreateReq{
			Code: consts.RoleCode_ROOT,
			//Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on name, this field is required", e.Error())
	})

	t.Run("should return duplicate code", func(t *testing.T) {
		req := CreateReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoCreate := roleadapter.RepoCreate{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}

		roleRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Role{}, adapter.ErrDuplicate)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "duplicate role code", e.Error())
	})

	t.Run("should return create query error", func(t *testing.T) {
		req := CreateReq{
			Code:            consts.RoleCode_ROOT,
			Name:            "Root",
			PermissionCodes: []types.Code{consts.PermissionCode_PermisisonRead},
		}
		repoCreate := roleadapter.RepoCreate{
			Code:            req.Code,
			Name:            req.Name,
			PermissionCodes: req.PermissionCodes,
		}

		roleRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Role{}, adapter.ErrQuery)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "create role query error", e.Error())
	})
}
