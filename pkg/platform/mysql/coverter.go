package mysql

import (
	"golang-projects-a/pkg/core/domain"
)

func (p Permission) load(perm domain.Permission) Permission {
	return Permission{
		ID:         perm.ID,
		Code:       perm.Code,
		Name:       perm.Name,
		Restricted: perm.Restricted,
	}
}

func (p Permission) domain() domain.Permission {
	return domain.Permission{
		ID:         p.ID,
		Code:       p.Code,
		Name:       p.Name,
		Restricted: p.Restricted,
	}
}

func (r Role) load(role domain.Role) Role {
	return Role{
		ID:   role.ID,
		Code: role.Code,
		Name: role.Name,
	}
}

func (r Role) domain() domain.Role {
	return domain.Role{
		ID:   r.ID,
		Code: r.Code,
		Name: r.Name,
		Permissions: func() []domain.Permission {
			perms := make([]domain.Permission, len(r.Permissions))
			for index, perm := range r.Permissions {
				perms[index] = perm.domain()
			}
			return perms
		}(),
	}
}

func (r RolePermission) load(rolePerm domain.RolePermission) RolePermission {
	return RolePermission{
		ID:           rolePerm.ID,
		RoleID:       rolePerm.RoleID,
		PermissionID: rolePerm.PermissionID,
	}
}

func (r RolePermission) domain() domain.RolePermission {
	return domain.RolePermission{
		ID:           r.ID,
		RoleID:       r.RoleID,
		PermissionID: r.PermissionID,
		Role: func() *domain.Role {
			if r.Role == nil {
				return nil
			}
			role := r.Role.domain()
			return &role
		}(),
		Permission: func() *domain.Permission {
			if r.Permission == nil {
				return nil
			}
			perm := r.Permission.domain()
			return &perm
		}(),
	}
}
