package user

import (
	"context"

	"twatter/pkg/core/adapter/useradapter"
	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
)

type UseCase interface {
	Find(ctx context.Context, req FindReq) (resp FindResp, serviceErr service.Error)
	GetList(ctx context.Context, req GetListReq) (resp GetListResp, serviceErr service.Error)

	Create(ctx context.Context, req CreateReq) (resp CreateResp, serviceErr service.Error)
	Update(ctx context.Context, req UpdateReq) (resp UpdateResp, serviceErr service.Error)
	Delete(ctx context.Context, req DeleteReq) (serviceErr service.Error)
}

type Service struct {
	UserRepo  useradapter.RepoAdapter
	Validator validatoradapter.Adapter
}

var _ UseCase = Service{}
