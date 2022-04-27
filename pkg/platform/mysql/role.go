package mysql

import (
	"context"

	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/roleadapter"
	"twatter/pkg/core/domain"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type roleRepo struct {
	db        *gorm.DB
	paginator paginator
}

var _ roleadapter.RepoAdapter = roleRepo{}

func (r roleRepo) Find(ctx context.Context, id uint32) (role domain.Role, err error) {
	roleModel := Role{}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return r.db.WithContext(ctx).Table(table_Roles).First(&roleModel, "id = ?", id).Error
	})
	g.Go(func() error {
		return r.db.WithContext(ctx).Table(table_Permissions).Find(&roleModel.Permissions).Error
	})

	err = g.Wait()
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return role, adapter.ErrNotFound
		default:
			return role, adapter.ErrQuery
		}
	}

	return roleModel.domain(), nil
}

func (r roleRepo) GetList(ctx context.Context, filter roleadapter.RepoFilter) (roles domain.Roles, err error) {
	roleModels := make([]Role, 0)
	rowCount := uint32(0)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return r.filterQuery(ctx, r.db, filter, true).Order("code ASC").Find(&roleModels).Error
	})
	g.Go(func() error {
		return r.filterQuery(ctx, r.db, filter, false).Select("COUNT(1)").Scan(&rowCount).Error
	})

	err = g.Wait()
	if err != nil {
		return roles, adapter.ErrQuery
	}

	return domain.Roles{
		Items: func() []domain.Role {
			items := make([]domain.Role, len(roleModels))
			for index, roleModel := range roleModels {
				items[index] = roleModel.domain()
			}
			return items
		}(),
		RowCount: rowCount,
	}, nil
}

func (r roleRepo) Create(ctx context.Context, data roleadapter.RepoCreate) (role domain.Role, err error) {
	roleModel := Role{}.load(domain.Role{
		Code: data.Code,
		Name: data.Name,
	})

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.WithContext(ctx).Table(table_Roles).Create(&roleModel).Error; err != nil {
			switch {
			case errorIs(err, errorCode_ER_DUP_ENTRY):
				return adapter.ErrDuplicate
			default:
				return adapter.ErrQuery
			}
		}

		permModels := make([]Permission, 0)
		if err = tx.WithContext(ctx).Table(table_Permissions).Where("code IN ?", data.PermissionCodes).Find(&permModels).Error; err != nil {
			return adapter.ErrQuery
		}
		roleModel.Permissions = permModels

		if len(permModels) > 0 {
			rolePermModels := make([]RolePermission, len(permModels))
			for index, permModel := range permModels {
				rolePermModels[index] = RolePermission{}.load(domain.RolePermission{
					RoleID:       roleModel.ID,
					PermissionID: permModel.ID,
				})
			}
			if err = tx.WithContext(ctx).Table(table_RolePermissions).Create(&rolePermModels).Error; err != nil {
				return adapter.ErrQuery
			}
		}

		return nil
	})
	if err != nil {
		return role, err
	}

	return roleModel.domain(), nil
}

func (r roleRepo) Update(ctx context.Context, id uint32, data roleadapter.RepoUpdate) (role domain.Role, err error) {
	roleModel := Role{}

	err = r.db.WithContext(ctx).Table(table_Roles).First(&roleModel).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return role, adapter.ErrNotFound
		default:
			return role, adapter.ErrQuery
		}
	}

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.WithContext(ctx).Table(table_RolePermissions).Where("role_id = ?", id).Unscoped().Delete(&RolePermission{}).Error; err != nil {
			return adapter.ErrQuery
		}

		// Update values
		if data.Name != "" {
			roleModel.Name = data.Name
		}

		if err = tx.WithContext(ctx).Table(table_Roles).Save(&roleModel).Error; err != nil {
			return adapter.ErrQuery
		}

		permModels := make([]Permission, 0)
		if err = tx.WithContext(ctx).Table(table_Permissions).Where("code IN ?", data.PermissionCodes).Find(&permModels).Error; err != nil {
			return adapter.ErrQuery
		}
		roleModel.Permissions = permModels

		if len(permModels) > 0 {
			rolePermModels := make([]RolePermission, len(permModels))
			for index, permModel := range permModels {
				rolePermModels[index] = RolePermission{}.load(domain.RolePermission{
					RoleID:       roleModel.ID,
					PermissionID: permModel.ID,
				})
			}
			if err = tx.WithContext(ctx).Table(table_RolePermissions).Create(&rolePermModels).Error; err != nil {
				return adapter.ErrQuery
			}
		}

		return nil
	})
	if err != nil {
		return role, err
	}

	return roleModel.domain(), nil
}

func (r roleRepo) Delete(ctx context.Context, id uint32) (err error) {
	var query *gorm.DB

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query = tx.WithContext(ctx).Table(table_RolePermissions).Where("role_id = ?", id).Unscoped().Delete(&RolePermission{})
		if err = query.Error; err != nil {
			return adapter.ErrQuery
		}

		query = tx.WithContext(ctx).Table(table_Roles).Unscoped().Delete(&Role{ID: id})
		if err = query.Error; err != nil {
			return adapter.ErrQuery
		}
		if query.RowsAffected == 0 {
			return adapter.ErrNotFound
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r roleRepo) filterQuery(ctx context.Context, db *gorm.DB, filter roleadapter.RepoFilter, pageEnabled bool) *gorm.DB {
	query := db.WithContext(ctx).Table(table_Roles)
	if filter.Code != "" {
		query = query.Where("code LIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filterPermCodes := filter.OmitemptyPermissionCodes(); len(filterPermCodes) > 0 {
		query = query.Where("id IN (?)", db.WithContext(ctx).Table(table_RolePermissions).Select("role_id").
			Where("permission_id IN (?)", db.WithContext(ctx).Table(table_Permissions).Select("id").
				Where("code IN ?", filterPermCodes)))
	}
	if pageEnabled && filter.PageSize > 0 {
		query = r.paginator.paginate(query, filter.PageNo, filter.PageSize)
	}

	return query
}
