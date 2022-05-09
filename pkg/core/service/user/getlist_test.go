package user

import (
	"context"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_useradapter "golang-projects-a/pkg/mocks/adapter/useradapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_useradapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		UserRepo:  userRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
			Limit:          10,
			Offset:         0,
		}
		repoFilter := useradapter.RepoFilter{
			ID:             req.ID,
			Username:       req.Username,
			Email:          req.Email,
			Password:       req.Password,
			AvatarUrl:      req.AvatarUrl,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
			Limit:          req.Limit,
			Offset:         req.Offset,
		}
		users := domain.Users{
			Items: []domain.User{
				{
					ID:             1,
					Username:       "Username",
					Email:          "Email",
					Password:       "Password",
					AvatarUrl:      "AvatarUrl",
					OrganizationId: 99,
					FollowingCount: 99,
					FollowerCount:  99,
				},
			},
			RowCount: 1,
		}

		userRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(users, nil)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{
			Items:    users.Items,
			RowCount: users.RowCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input limit greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
			Limit:          -10,
			Offset:         0,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on limit, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return invalid input offset greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
			Limit:          10,
			Offset:         -1,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on offset, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return get list query error", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
			Limit:          10,
			Offset:         0,
		}
		repoFilter := useradapter.RepoFilter{
			ID:             req.ID,
			Username:       req.Username,
			Email:          req.Email,
			Password:       req.Password,
			AvatarUrl:      req.AvatarUrl,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
			Limit:          req.Limit,
			Offset:         req.Offset,
		}

		userRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(domain.Users{}, adapter.ErrQuery)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "get list user query error", e.Error())
	})
}
