package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_useradapter "golang-projects-a/pkg/mocks/adapter/useradapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_useradapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		UserRepo:  userRepo,
		Validator: v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
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

		userRepo.EXPECT().Find(gomock.Any(), req.id).Return(user, nil)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{
			ID:             user.ID,
			Username:       user.Username,
			Email:          user.Email,
			Password:       user.Password,
			AvatarUrl:      user.AvatarUrl,
			OrganizationId: user.OrganizationId,
			FollowingCount: user.FollowingCount,
			FollowerCount:  user.FollowerCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := FindReq{
			id: 1,
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

	t.Run("should return user not found", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		userRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.User{}, adapter.ErrNotFound)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "user not found", e.Error())
	})

	t.Run("should return find query error", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		userRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.User{}, adapter.ErrQuery)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "find user query error", e.Error())
	})
}
