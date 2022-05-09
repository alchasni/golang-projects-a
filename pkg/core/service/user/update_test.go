package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_useradapter "golang-projects-a/pkg/mocks/adapter/useradapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_useradapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		UserRepo:  userRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,

			id: 1,
		}
		repoUpdate := useradapter.RepoUpdate{
			Username:       req.Username,
			Email:          req.Email,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}
		user := domain.User{
			ID:             1,
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		userRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(user, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:             1,
			Username:       user.Username,
			Email:          user.Email,
			OrganizationId: user.OrganizationId,
			FollowingCount: user.FollowingCount,
			FollowerCount:  user.FollowerCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow no updating", func(t *testing.T) {
		req := UpdateReq{
			ID: "1",

			id: 1,
		}
		repoUpdate := useradapter.RepoUpdate{
			Username:       req.Username,
			Email:          req.Email,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}
		user := domain.User{
			ID:             1,
			Username:       "new username",
			Email:          "new email",
			AvatarUrl:      "avatarUrl",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,
		}

		userRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(user, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:             1,
			Username:       user.Username,
			Email:          user.Email,
			OrganizationId: user.OrganizationId,
			FollowingCount: user.FollowingCount,
			FollowerCount:  user.FollowerCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := UpdateReq{
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,

			id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := UpdateReq{
			ID:             "a",
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,

			id: 0,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return user not found", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,

			id: 1,
		}
		repoUpdate := useradapter.RepoUpdate{
			Username:       req.Username,
			Email:          req.Email,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}

		userRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.User{}, adapter.ErrNotFound)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "user not found", e.Error())
	})

	t.Run("should return update query error", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Username:       "new username",
			Email:          "new email",
			OrganizationId: 99,
			FollowingCount: 99,
			FollowerCount:  99,

			id: 1,
		}
		repoUpdate := useradapter.RepoUpdate{
			Username:       req.Username,
			Email:          req.Email,
			OrganizationId: req.OrganizationId,
			FollowingCount: req.FollowingCount,
			FollowerCount:  req.FollowerCount,
		}

		userRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.User{}, adapter.ErrQuery)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "update user query error", e.Error())
	})
}
