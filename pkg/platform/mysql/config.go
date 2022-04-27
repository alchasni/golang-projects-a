package mysql

import (
	"fmt"
	"net/url"

	"golang-projects-a/pkg/consts"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB         DBConfig         `yaml:"db"`
	Pagination PaginationConfig `yaml:"pagination"`
}

type PaginationConfig struct {
	MinPageSize int `yaml:"min_page_size" validate:"gte=0"`
	MaxPageSize int `yaml:"max_page_size" validate:"gte=0"`
}

type DBConfig struct {
	User                     string `yaml:"user" validate:"required"`
	Pass                     string `yaml:"pass" validate:"required" logger:"-"`
	Host                     string `yaml:"host" validate:"required"`
	Port                     string `yaml:"port" validate:"required"`
	DB                       string `yaml:"db" validate:"required"`
	MaxIdleConns             int    `yaml:"max_idle_conns"`
	MaxOpenConns             int    `yaml:"max_open_conns"`
	SlowQueryThresholdMillis int    `yaml:"slow_query_threshold_millis"`
}

func (c DBConfig) ConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		c.User, c.Pass, c.Host, c.Port, c.DB, url.QueryEscape(consts.Location_AsiaJakarta))
}

func initORM(cfg DBConfig) (*gorm.DB, error) {
	orm, err := gorm.Open(mysql.Open(cfg.ConnStr()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	// Table migration
	err = orm.AutoMigrate(&Role{}, &Permission{}, &RolePermission{})
	if err != nil {
		return nil, err
	}

	db, err := orm.DB()
	if err == nil {
		if cfg.MaxIdleConns > 0 {
			db.SetMaxIdleConns(cfg.MaxIdleConns)
		}
		if cfg.MaxOpenConns > 0 {
			db.SetMaxOpenConns(cfg.MaxOpenConns)
		}
	}

	return orm, nil
}
