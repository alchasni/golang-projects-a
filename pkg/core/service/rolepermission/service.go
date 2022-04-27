package rolepermission

import (
	"golang-projects-a/pkg/core/adapter/rolepermissionadapter"
	"golang-projects-a/pkg/core/adapter/validatoradapter"
)

type UseCase interface {
}

type Service struct {
	RolePermissionRepo rolepermissionadapter.RepoAdapter
	Validator          validatoradapter.Adapter
}

var _ UseCase = Service{}
