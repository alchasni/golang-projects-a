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

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mock_commentadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		CommentRepo: commentRepo,
		Validator:   v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Comment:        "Comment",
			OrganizationId: 1,

			id: 1,
		}
		repoUpdate := commentadapter.RepoUpdate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}
		comment := domain.Comment{
			ID:             1,
			Comment:        "Comment",
			OrganizationId: 1,
		}

		commentRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(comment, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:             comment.ID,
			Comment:        comment.Comment,
			OrganizationId: comment.OrganizationId,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow no updating", func(t *testing.T) {
		req := UpdateReq{
			ID: "1",

			id: 1,
		}
		repoUpdate := commentadapter.RepoUpdate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}
		comment := domain.Comment{
			ID:             1,
			Comment:        "Comment",
			OrganizationId: 1,
		}

		commentRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(comment, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:             comment.ID,
			Comment:        comment.Comment,
			OrganizationId: comment.OrganizationId,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := UpdateReq{
			Comment:        "Comment",
			OrganizationId: 1,

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
			Comment:        "Comment",
			OrganizationId: 1,

			id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return comment not found", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Comment:        "Comment",
			OrganizationId: 1,

			id: 1,
		}
		repoUpdate := commentadapter.RepoUpdate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}

		commentRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Comment{}, adapter.ErrNotFound)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "comment not found", e.Error())
	})

	t.Run("should return update query error", func(t *testing.T) {
		req := UpdateReq{
			ID:             "1",
			Comment:        "Comment",
			OrganizationId: 1,

			id: 1,
		}
		repoUpdate := commentadapter.RepoUpdate{
			Comment:        req.Comment,
			OrganizationId: req.OrganizationId,
		}

		commentRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Comment{}, adapter.ErrQuery)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "update comment query error", e.Error())
	})
}
