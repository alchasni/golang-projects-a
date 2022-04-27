package auth

import (
	"context"

	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type UpdatePasswordReq struct{}

type UpdatePasswordResp struct{}

func (req UpdatePasswordReq) validate(v validatoradapter.Adapter) error {
	panic("implement me")
}

func (s Service) UpdatePassword(ctx context.Context, req UpdatePasswordReq) (serviceErr service.Error) {
	panic("implement me")
}
