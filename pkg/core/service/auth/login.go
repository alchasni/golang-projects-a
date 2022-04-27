package auth

import (
	"context"

	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
)

type LoginReq struct{}

type LoginResp struct{}

func (req LoginReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) Login(ctx context.Context, req LoginReq) (resp LoginResp, serviceErr service.Error) {
	panic("implement me")
}
