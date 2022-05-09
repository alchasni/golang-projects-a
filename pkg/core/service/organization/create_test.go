package organization

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_organizationadapter "golang-projects-a/pkg/mocks/adapter/organizationadapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRepo := mock_organizationadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		OrganizationRepo: organizationRepo,
		Validator:        v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := CreateReq{
			Name: "Name",
		}
		repoCreate := organizationadapter.RepoCreate{
			Name: req.Name,
		}
		org := domain.Organization{
			ID:   1,
			Name: "Name",
		}

		organizationRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(org, nil)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{
			ID:   org.ID,
			Name: org.Name,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input name required", func(t *testing.T) {
		req := CreateReq{}

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on name, this field should only contains alphanumeric value", e.Error())
	})

	t.Run("should return create query error", func(t *testing.T) {
		req := CreateReq{
			Name: "Name",
		}
		repoCreate := organizationadapter.RepoCreate{
			Name: req.Name,
		}

		organizationRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Organization{}, adapter.ErrQuery)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "create organization query error", e.Error())
	})

	t.Run("should return duplicate error", func(t *testing.T) {
		req := CreateReq{
			Name: "Name",
		}
		repoCreate := organizationadapter.RepoCreate{
			Name: req.Name,
		}

		organizationRepo.EXPECT().Create(gomock.Any(), repoCreate).Return(domain.Organization{}, adapter.ErrDuplicate)

		resp, e := svc.Create(context.Background(), req)
		assert.Equal(t, CreateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "duplicate organization name", e.Error())
	})

}
