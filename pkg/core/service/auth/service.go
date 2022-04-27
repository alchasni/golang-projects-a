package auth

import (
	"context"

	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
	"golang-projects-a/pkg/core/service"
)

type UseCase interface {
	Login(ctx context.Context, req LoginReq) (resp LoginResp, serviceErr service.Error)
	Logout(ctx context.Context, req LogoutReq) (serviceErr service.Error)
	UpdatePassword(ctx context.Context, req UpdatePasswordReq) (serviceErr service.Error)
}

type Service struct {
	UserRepo  useradapter.RepoAdapter
	Validator validatoradapter.Adapter
}

var _ UseCase = Service{}
