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

func TestService_FindByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRepo := mock_organizationadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		OrganizationRepo: organizationRepo,
		Validator:        v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := FindByNameReq{
			Name: "Name",
		}
		organization := domain.Organization{
			ID:   1,
			Name: "Name",
		}

		organizationRepo.EXPECT().FindByName(gomock.Any(), req.Name).Return(organization, nil)

		resp, e := svc.FindByName(context.Background(), req)
		assert.Equal(t, FindByNameResp{
			ID: organization.ID,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input name required", func(t *testing.T) {
		req := FindByNameReq{}

		resp, e := svc.FindByName(context.Background(), req)
		assert.Equal(t, FindByNameResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on name, this field is required", e.Error())
	})

	t.Run("should return organization not found", func(t *testing.T) {
		req := FindByNameReq{
			Name: "Name",
		}

		organizationRepo.EXPECT().FindByName(gomock.Any(), req.Name).Return(domain.Organization{}, adapter.ErrNotFound)

		resp, e := svc.FindByName(context.Background(), req)
		assert.Equal(t, FindByNameResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "organization not found", e.Error())
	})

	t.Run("should return find query error", func(t *testing.T) {
		req := FindByNameReq{
			Name: "Name",
		}

		organizationRepo.EXPECT().FindByName(gomock.Any(), req.Name).Return(domain.Organization{}, adapter.ErrQuery)

		resp, e := svc.FindByName(context.Background(), req)
		assert.Equal(t, FindByNameResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "find organization query error", e.Error())
	})
}
