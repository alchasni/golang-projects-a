package organization

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/domain"
	"golang-projects-a/pkg/core/service"
	mock_organizationadapter "golang-projects-a/pkg/mocks/adapter/organizationadapter"
	"golang-projects-a/pkg/platform/validator"
	"testing"
)

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRepo := mock_organizationadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		OrganizationRepo: organizationRepo,
		Validator:        v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}
		organization := domain.Organization{
			ID:   1,
			Name: "Name",
		}

		organizationRepo.EXPECT().Find(gomock.Any(), req.id).Return(organization, nil)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{
			ID:   organization.ID,
			Name: organization.Name,
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

	t.Run("should return organization not found", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		organizationRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Organization{}, adapter.ErrNotFound)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "organization not found", e.Error())
	})

	t.Run("should return find query error", func(t *testing.T) {
		req := FindReq{
			ID: "1",

			id: 1,
		}

		organizationRepo.EXPECT().Find(gomock.Any(), req.id).Return(domain.Organization{}, adapter.ErrQuery)

		resp, e := svc.Find(context.Background(), req)
		assert.Equal(t, FindResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "find organization query error", e.Error())
	})
}
