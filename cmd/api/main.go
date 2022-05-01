package main

import (
	"flag"
	"fmt"
	"golang-projects-a/pkg/core/service/user"
	"golang-projects-a/pkg/platform/mongodb"

	"golang-projects-a/pkg/platform/validator"
	"golang-projects-a/pkg/platform/yaml"
	"golang-projects-a/pkg/transport/http"
)

func main() {
	v := validator.New()

	cfgPath := flag.String("configpath", "cmd/api/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := yaml.Init(*cfgPath, v)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	//mysqlDB, err := mysql.New(cfg.Datasource.MySQL)
	//if err != nil {
	//	panic(fmt.Errorf("error mysql initialization. %w", err))
	//}

	mongoDB, err := mongodb.New(cfg.Datasource.MongoDB)
	if err != nil {
		panic(fmt.Errorf("error mysql initialization. %w", err))
	}

	userService := user.Service{
		UserRepo:  mongoDB.UserRepo(),
		Validator: v,
	}

	httpServer := http.HTTP{
		UserService: userService,
		Env:         cfg.Server.Env,
		Config:      cfg.Server.HTTP,
	}
	if err = httpServer.Init(); err != nil {
		panic(fmt.Errorf("error http server initialization. %w", err))
	}

	httpServer.Serve()
}
