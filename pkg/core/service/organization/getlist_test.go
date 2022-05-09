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

func TestService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRepo := mock_organizationadapter.NewMockRepoAdapter(ctrl)
	v := validator.New()

	svc := Service{
		OrganizationRepo: organizationRepo,
		Validator:        v,
	}

	t.Run("happy flow", func(t *testing.T) {
		req := GetListReq{
			ID:     1,
			Name:   "Name",
			Limit:  10,
			Offset: 0,
		}
		repoFilter := organizationadapter.RepoFilter{
			ID:     req.ID,
			Name:   req.Name,
			Limit:  req.Limit,
			Offset: req.Offset,
		}
		organizations := domain.Organizations{
			Items: []domain.Organization{
				{
					ID:   1,
					Name: "Name",
				},
			},
			RowCount: 1,
		}

		organizationRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(organizations, nil)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{
			Items:    organizations.Items,
			RowCount: organizations.RowCount,
		}, resp)
		assert.Equal(t, nil, e)
	})

	t.Run("should return invalid input limit greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			ID:     1,
			Name:   "Name",
			Limit:  -10,
			Offset: 0,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on limit, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return invalid input offset greater than or equal 0", func(t *testing.T) {
		req := GetListReq{
			ID:     1,
			Name:   "Name",
			Limit:  10,
			Offset: -1,
		}

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_InvalidInput, e.Code())
		assert.Equal(t, "invalid input on offset, this field should be greater than or equal 0", e.Error())
	})

	t.Run("should return get list query error", func(t *testing.T) {
		req := GetListReq{
			ID:     1,
			Name:   "Name",
			Limit:  10,
			Offset: 0,
		}
		repoFilter := organizationadapter.RepoFilter{
			ID:     req.ID,
			Name:   req.Name,
			Limit:  req.Limit,
			Offset: req.Offset,
		}

		organizationRepo.EXPECT().GetList(gomock.Any(), repoFilter).Return(domain.Organizations{}, adapter.ErrQuery)

		resp, e := svc.GetList(context.Background(), req)
		assert.Equal(t, GetListResp{}, resp)
		assert.Equal(t, service.ErrorCode_DatasourceAccess, e.Code())
		assert.Equal(t, "get list organization query error", e.Error())
	})
}
