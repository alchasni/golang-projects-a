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
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}
		repoCreate := useradapter.RepoCreate{
			Username:       req.Username,
			Email:          req.Email,
			Password:       req.Password,
			AvatarUrl:      req.AvatarUrl,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}
		user := domain.User{
			ID:             1,
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		userRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(user, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:             user.ID,
			Username:       user.Username,
			Email:          user.Email,
			AvatarUrl:      user.AvatarUrl,
			OrganizationId: user.OrganizationId,
			FollowingCount: user.FollowingCount,
			FollowerCount:  user.FollowerCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input username required", func(t *testing.T) {
		req := CreateReq{
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on username, this field is required", e.Error())
	})

	t.Run("should return invalid input email required", func(t *testing.T) {
		req := CreateReq{
			Username:       "Username",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on email, this field is required", e.Error())
	})

	t.Run("should return invalid input password required", func(t *testing.T) {
		req := CreateReq{
			Username:       "Username",
			Email:          "Email",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on password, this field is required", e.Error())
	})

	t.Run("should return invalid input avatar_url required", func(t *testing.T) {
		req := CreateReq{
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on avatar_url, this field is required", e.Error())
	})

	t.Run("should return invalid input organization_id required", func(t *testing.T) {
		req := CreateReq{
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			FollowingCount: 99,
			FollowerCount:  99,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on organization_id, this field is required", e.Error())
	})

	t.Run("should return create query error", func(t *testing.T) {
		req := CreateReq{
			Username:       "Username",
			Email:          "Email",
			Password:       "Password",
			AvatarUrl:      "AvatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}
		repoCreate := useradapter.RepoCreate{
			Username:       req.Username,
			Email:          req.Email,
			Password:       req.Password,
			AvatarUrl:      req.AvatarUrl,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}

		userRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.User{}, adapter.ErrQuery)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "create user query error", e.Error())
	})
}
