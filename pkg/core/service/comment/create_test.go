package comment

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_commentadapter "golang-projects-a/pkg/mocks/adapter/commentadapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mock_commentadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		CommentRepo: commentRepo,
		Validator:   v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CreateReq{
			Comment:        "Comment",
			OrganizationId: 1,
		}
		repoCreate := commentadapter.RepoCreate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}
		org := domain.Comment{
			ID:             1,
			Comment:        "Comment",
			OrganizationId: 1,
		}

		commentRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(org, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:             org.ID,
			Comment:        org.Comment,
			OrganizationId: org.OrganizationId,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input comment required", func(t *testing.T) {
		req := CreateReq{
			OrganizationId: 1,
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on comment, this field is required", e.Error())
	})

	t.Run("should return invalid input organization_id required", func(t *testing.T) {
		req := CreateReq{
			Comment: "Comment",
		}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on organization_id, this field is required", e.Error())
	})

	t.Run("should return create query error", func(t *testing.T) {
		req := CreateReq{
			Comment:        "Comment",
			OrganizationId: 1,
		}
		repoCreate := commentadapter.RepoCreate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}

		commentRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Comment{}, adapter.ErrQuery)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "create comment query error", e.Error())
	})
}
