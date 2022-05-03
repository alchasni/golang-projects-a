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

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRepo := mock_organizationadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		OrganizationRepo: organizationRepo,
		Validator:        v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := UpdateReq{
			ID:   "1",
			Name: "Name",

			id: 1,
		}
		repoUpdate := organizationadapter.RepoUpdate{
			Name: req.Name,
		}
		organization := domain.Organization{
			ID:   1,
			Name: "Name",
		}

		organizationRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(organization, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:   organization.ID,
			Name: organization.Name,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("happy flow no updating", func(t *testing.T) {
		req := UpdateReq{
			ID: "1",

			id: 1,
		}
		repoUpdate := organizationadapter.RepoUpdate{
			Name: req.Name,
		}
		organization := domain.Organization{
			ID:   1,
			Name: "Name",
		}

		organizationRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(organization, nil)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{
			ID:   organization.ID,
			Name: organization.Name,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input id required", func(t *testing.T) {
		req := UpdateReq{
			Name: "Name",

			id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field is required", e.Error())
	})

	t.Run("should return invalid input id numeric", func(t *testing.T) {
		req := UpdateReq{
			ID:   "a",
			Name: "Name",

			id: 1,
		}

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on id, this field should only contains numeric value", e.Error())
	})

	t.Run("should return organization not found", func(t *testing.T) {
		req := UpdateReq{
			ID:   "1",
			Name: "Name",

			id: 1,
		}
		repoUpdate := organizationadapter.RepoUpdate{
			Name: req.Name,
		}

		organizationRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Organization{}, adapter.ErrNotFound)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "organization not found", e.Error())
	})

	t.Run("should return update query error", func(t *testing.T) {
		req := UpdateReq{
			ID:   "1",
			Name: "Name",

			id: 1,
		}
		repoUpdate := organizationadapter.RepoUpdate{
			Name: req.Name,
		}

		organizationRepo.EXPECT().Update(gomock.Any(), req.id, repoUpdate).Return(domain.Organization{}, adapter.ErrQuery)

		resp, e := svc.Update(context.Background(), req)
		assert.Equal(t, UpdateResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "update organization query error", e.Error())
	})
}
