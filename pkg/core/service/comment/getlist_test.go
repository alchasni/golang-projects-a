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

func TestService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mock_commentadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		CommentRepo: commentRepo,
		Validator:   v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			OrganizationId: 1,
			Limit:          10,
			Offset:         0,
		}
		repoFilter := commentadapter.RepoFilter{
			ID:             req.ID,
			OrganizationId: req.OrganizationId,
			Limit:          req.Limit,
			Offset:         req.Offset,
		}
		comments := domain.Comments{
			Items: []domain.Comment{
				{
					ID:             1,
					Comment:        "Comment",
					OrganizationId: 1,
				},
			},
			RowCount: 1,
		}

		commentRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(comments, nil)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{
			Items:    comments.Items,
			RowCount: comments.RowCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input limit greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			ID:             1,
			OrganizationId: 1,
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
			OrganizationId: 1,
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
			OrganizationId: 1,
			Limit:          10,
			Offset:         0,
		}
		repoFilter := commentadapter.RepoFilter{
			ID:             req.ID,
			OrganizationId: req.OrganizationId,
			Limit:          req.Limit,
			Offset:         req.Offset,
		}

		commentRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(domain.Comments{}, adapter.ErrQuery)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "get list comment query error", e.Error())
	})
}
