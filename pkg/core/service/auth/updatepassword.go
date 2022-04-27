package auth

import (
	"context"

	"twatter/pkg/core/adapter/validatoradapter"
	"twatter/pkg/core/service"
)

type UpdatePasswordReq struct{}

type UpdatePasswordResp struct{}

func (req UpdatePasswordReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) UpdatePassword(ctx context.Context, req UpdatePasswordReq) (serviceErr service.Error) {
	panic("implement me")
}
