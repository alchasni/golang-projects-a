package mysql

import (
	"context"

	"golang.org/x/sync/errgroup"
	"twatter/pkg/core/adapter"
	"twatter/pkg/core/adapter/permissionadapter"
	"twatter/pkg/core/domain"

	"gorm.io/gorm"
)

type permissionRepo struct {
	db        *gorm.DB
	paginator paginator
}

var _ permissionadapter.RepoAdapter = permissionRepo{}

func (p permissionRepo) Find(ctx context.Context, id uint32) (permission domain.Permission, err error) {
	permModel := Permission{}

	err = p.db.WithContext(ctx).Table(table_Permissions).First(&permModel, "id = ?", id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return permission, adapter.ErrNotFound
		default:
			return permission, adapter.ErrQuery
		}
	}

	return permModel.domain(), nil
}

func (p permissionRepo) GetList(ctx context.Context, filter permissionadapter.RepoFilter) (permissions domain.Permissions, err error) {
	permModels := make([]Permission, 0)
	rowCount := uint32(0)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return p.filterQuery(ctx, p.db, filter, true).Order("code ASC").Find(&permModels).Error
	})
	g.Go(func() error {
		return p.filterQuery(ctx, p.db, filter, false).Select("COUNT(1)").Scan(&rowCount).Error
	})

	err = g.Wait()
	if err != nil {
		return permissions, adapter.ErrQuery
	}

	return domain.Permissions{
		Items: func() []domain.Permission {
			items := make([]domain.Permission, len(permModels))
			for index, permModel := range permModels {
				items[index] = permModel.domain()
			}
			return items
		}(),
		RowCount: rowCount,
	}, nil
}

func (p permissionRepo) Create(ctx context.Context, data permissionadapter.RepoCreate) (permission domain.Permission, err error) {
	permModel := Permission{}.load(domain.Permission{
		Code:       data.Code,
		Name:       data.Name,
		Restricted: data.Restricted,
	})

	err = p.db.WithContext(ctx).Table(table_Permissions).Create(&permModel).Error
	if err != nil {
		switch {
		case errorIs(err, errorCode_ER_DUP_ENTRY):
			return permission, adapter.ErrDuplicate
		default:
			return permission, adapter.ErrQuery
		}
	}

	return permModel.domain(), nil
}

func (p permissionRepo) Update(ctx context.Context, id uint32, data permissionadapter.RepoUpdate) (permission domain.Permission, err error) {
	permModel := Permission{}

	err = p.db.WithContext(ctx).Table(table_Permissions).First(&permModel, "id = ?", id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return permission, adapter.ErrNotFound
		default:
			return permission, adapter.ErrQuery
		}
	}

	// Update values
	if data.Name != "" {
		permModel.Name = data.Name
	}
	if data.Restricted != "" {
		permModel.Restricted = data.Restricted
	}

	err = p.db.WithContext(ctx).Table(table_Permissions).Save(&permModel).Error
	if err != nil {
		return permission, adapter.ErrQuery
	}

	return permModel.domain(), nil
}

func (p permissionRepo) Delete(ctx context.Context, id uint32) (err error) {
	query := p.db.WithContext(ctx).Table(table_Permissions).Unscoped().Delete(&Permission{ID: id})

	err = query.Error
	if err != nil {
		return adapter.ErrQuery
	}
	if query.RowsAffected == 0 {
		return adapter.ErrNotFound
	}

	return nil
}

func (p permissionRepo) filterQuery(ctx context.Context, db *gorm.DB, filter permissionadapter.RepoFilter, pageEnabled bool) *gorm.DB {
	query := db.WithContext(ctx).Table(table_Permissions)
	if filter.Code != "" {
		query = query.Where("code LIKE ?", "%"+filter.Code+"%")
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.Restricted != "" {
		query = query.Where("restricted = ?", filter.Restricted)
	}
	if pageEnabled && filter.PageSize > 0 {
		query = p.paginator.paginate(query, filter.PageNo, filter.PageSize)
	}

	return query
}
