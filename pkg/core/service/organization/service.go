package organization

import (
	"context"
	"golang-projects-a/pkg/core/adapter/organizationadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type UseCase interface {
	Find(ctx context.Context, req FindReq) (resp FindResp, serviceErr service.Error)
	GetList(ctx context.Context, req GetListReq) (resp GetListResp, serviceErr service.Error)

	Create(ctx context.Context, req CreateReq) (resp CreateResp, serviceErr service.Error)
	Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error)
	Delete(ctx context.Context, req DeleteReq) (serviceErr service.Error)

	FindByName(ctx context.Context, req FindByNameReq) (resp FindByNameResp, serviceErr service.Error)
}

type Service struct {
	OrganizationRepo organizationadapter.RepoAdapter
	Validator        validatoradapter.Adapter
}

var _ UseCase = Service{}
