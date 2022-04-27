package main

import (
	"flag"
	"fmt"

	"twatter/pkg/core/service/permission"
	"twatter/pkg/core/service/role"
	"twatter/pkg/core/service/rolepermission"
	"twatter/pkg/platform/mysql"
	"twatter/pkg/platform/validator"
	"twatter/pkg/platform/yaml"
	"twatter/pkg/transport/http"
)

func main() {
	v := validator.New()

	cfgPath := flag.String("configpath", "cmd/api/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := yaml.Init(*cfgPath, v)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	mysqlDB, err := mysql.New(cfg.Datasource.MySQL)
	if err != nil {
		panic(fmt.Errorf("error mysql initialization. %w", err))
	}

	roleService := role.Service{
		RoleRepo:  mysqlDB.RoleRepo(),
		Validator: v,
	}

	permissionService := permission.Service{
		PermissionRepo: mysqlDB.PermissionRepo(),
		Validator:      v,
	}

	rolePermissionService := rolepermission.Service{
		RolePermissionRepo: mysqlDB.RolePermissionRepo(),
		Validator:          v,
	}

	httpServer := http.HTTP{
		PermissionService:     permissionService,
		RoleService:           roleService,
		RolePermissionService: rolePermissionService,
		Env:                   cfg.Server.Env,
		Config:                cfg.Server.HTTP,
	}
	if err = httpServer.Init(); err != nil {
		panic(fmt.Errorf("error http server initialization. %w", err))
	}

	httpServer.Serve()
}
