package user

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/consts"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"
	mock_useradapter "golang-projects-a/pkg/mocks/adapter/useradapter"
	"golang-projects-a/pkg/platform/validator"

	"github.com/golang/mock/gomock"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_useradapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		UserRepo:  userRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CreateReq{
			Username: "root",
			Email:    "root@example.com",
			Password: "password",
			RoleCode: consts.RoleCode_ROOT,
		}
		repoCreate := useradapter.RepoCreate{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			RoleCode: req.RoleCode,
		}
		user := domain.User{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  repoCreate.Username,
			Email:     repoCreate.Email,
			RoleCode:  repoCreate.RoleCode,
			Role: &domain.Role{
				ID:   1,
				Code: repoCreate.RoleCode,
				Name: "Root",
				Permissions: []domain.Permission{
					{
						ID:         1,
						Code:       consts.PermissionCode_PermisisonRead,
						Name:       "Permission Read",
						Restricted: consts.Confirmation_Y,
					},
				},
			},
		}

		userRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(user, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow omitempty", func(t *testing.T) {
		req := CreateReq{
			Username: "root",
			//Email:    "root@example.com",
			Password: "password",
			RoleCode: consts.RoleCode_ROOT,
		}
		repoCreate := useradapter.RepoCreate{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			RoleCode: req.RoleCode,
		}
		user := domain.User{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  repoCreate.Username,
			Email:     repoCreate.Email,
			RoleCode:  repoCreate.RoleCode,
			Role: &domain.Role{
				ID:   1,
				Code: repoCreate.RoleCode,
				Name: "Root",
				Permissions: []domain.Permission{
					{
						ID:         1,
						Code:       consts.PermissionCode_PermisisonRead,
						Name:       "Permission Read",
						Restricted: consts.Confirmation_Y,
					},
				},
			},
		}

		userRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(user, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
		}, resp)
		assert.Equal(t, nil, e)
	})
}
