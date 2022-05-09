package auth

import (
	"context"

	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type LogoutReq struct{}

type LogoutResp struct{}

func (req LogoutReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) Logout(ctx context.Context, req LogoutReq) (serviceErr service.Error) {
	panic("implement me")
}
