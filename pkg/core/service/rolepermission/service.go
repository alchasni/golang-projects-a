package rolepermission

import (
	"twatter/pkg/core/adapter/rolepermissionadapter"
	"twatter/pkg/core/adapter/validatoradapter"
)

type UseCase interface {
}

type Service struct {
	RolePermissionRepo rolepermissionadapter.RepoAdapter
	Validator          validatoradapter.Adapter
}

var _ UseCase = Service{}
