package mysql

import (
	"context"
	"golang-projects-a/pkg/core/adapter"
	"golang-projects-a/pkg/core/adapter/useradapter"
	"golang-projects-a/pkg/core/domain"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type userRepo struct {
	db        *gorm.DB
	paginator paginator
}

var _ useradapter.RepoAdapter = userRepo{}

func (u userRepo) Find(ctx context.Context, id uint64) (domain.User, error) {
	user := User{}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return u.db.WithContext(ctx).Table(tableUsers).First(&user, "id = ?", id).Error
	})

	err := g.Wait()
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return user.domain(), adapter.ErrNotFound
		default:
			return user.domain(), adapter.ErrQuery
		}
	}

	return user.domain(), nil
}

func (u userRepo) GetList(ctx context.Context, filter useradapter.RepoFilter) (domain.Users, error) {
	var res domain.Users
	rowCount := uint64(0)
	user := make([]User, rowCount)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return u.filterQuery(ctx, u.db, filter, true).Order("code ASC").Find(&user).Error
	})
	g.Go(func() error {
		return u.filterQuery(ctx, u.db, filter, false).Select("COUNT(1)").Scan(&rowCount).Error
	})

	err := g.Wait()
	if err != nil {
		return res, adapter.ErrQuery
	}

	return domain.Users{
		Items: func() []domain.User {
			items := make([]domain.User, len(user))
			for index, roleModel := range user {
				items[index] = roleModel.domain()
			}
			return items
		}(),
		RowCount: rowCount,
	}, nil
}

func (u userRepo) Create(ctx context.Context, data useradapter.RepoCreate) (domain.User, error) {
	user := User{}.load(domain.User{
		Username:  data.Username,
		Password:  data.Password,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
	})

	err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Table(tableUsers).Create(&user).Error; err != nil {
			switch {
			case errorIs(err, errorCode_ER_DUP_ENTRY):
				return adapter.ErrDuplicate
			default:
				return adapter.ErrQuery
			}
		}

		return nil
	})
	if err != nil {
		return user.domain(), err
	}

	return user.domain(), nil
}

func (u userRepo) Update(ctx context.Context, id uint64, data useradapter.RepoUpdate) (domain.User, error) {
	user := User{}

	err := u.db.WithContext(ctx).Table(tableUsers).First(&user).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return user.domain(), adapter.ErrNotFound
		default:
			return user.domain(), adapter.ErrQuery
		}
	}

	err = u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if data.Username != "" {
			user.Username = data.Username
		}

		if err = tx.WithContext(ctx).Table(tableUsers).Save(&user).Error; err != nil {
			return adapter.ErrQuery
		}

		return nil
	})
	if err != nil {
		return user.domain(), err
	}

	return user.domain(), nil
}

func (u userRepo) Delete(ctx context.Context, id uint64) (err error) {
	var query *gorm.DB

	err = u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query = tx.WithContext(ctx).Table(tableUsers).Unscoped().Delete(&User{ID: id})
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

func (u userRepo) filterQuery(ctx context.Context, db *gorm.DB, filter useradapter.RepoFilter, pageEnabled bool) *gorm.DB {
	query := db.WithContext(ctx).Table(tableUsers)
	if filter.Username != "" {
		query = query.Where("code LIKE ?", "%"+filter.Username+"%")
	}
	if filter.Email != "" {
		query = query.Where("name LIKE ?", "%"+filter.Email+"%")
	}
	if pageEnabled && filter.Limit > 0 {
		query = u.paginator.paginate(query, filter.Offset, filter.Limit)
	}

	return query
}
