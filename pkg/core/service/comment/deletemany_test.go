package comment

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/service"
	mock_commentadapter "golang-projects-a/pkg/mocks/adapter/commentadapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_DeleteMany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mock_commentadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		CommentRepo: commentRepo,
		Validator:   v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := DeleteManyReq{
			OrganizationId: 1,
		}
		repoFilter := commentadapter.RepoFilter{
			OrganizationId: req.OrganizationId,
		}

		commentRepo.EXPECT().DeleteMany(gomock.Any(), repoFilter).Return(nil)

		e := svc.DeleteMany(context.Background(), req)
		assert.Equal(t, nil, e)
	})

	t.Run("should return get list query error", func(t *testing.T) {
		req := DeleteManyReq{
			OrganizationId: 1,
		}
		repoFilter := commentadapter.RepoFilter{
			OrganizationId: req.OrganizationId,
		}

		commentRepo.EXPECT().DeleteMany(gomock.Any(), repoFilter).Return(adapter.ErrQuery)

		e := svc.DeleteMany(context.Background(), req)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "delete list comment query error", e.Error())
	})
}
