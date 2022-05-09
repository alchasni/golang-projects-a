package auth

import (
	"context"

	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type LoginReq struct{}

type LoginResp struct{}

func (req LoginReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) Login(ctx context.Context, req LoginReq) (resp LoginResp, serviceErr service.Error) {
	panic("implement me")
}
