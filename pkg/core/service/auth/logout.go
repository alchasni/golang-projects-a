package auth

import (
	"context"

	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
)

type LogoutReq struct{}

type LogoutResp struct{}

func (req LogoutReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) Logout(ctx context.Context, req LogoutReq) (serviceErr service.Error) {
	panic("implement me")
}
