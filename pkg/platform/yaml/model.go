package yaml

import (
	"golang-projects-a/pkg/platform/mongodb"
	"golang-projects-a/pkg/platform/mysql"
	"golang-projects-a/pkg/transport/http"
)

type Config struct {
	Server     server     `yaml:"server"`
	Datasource datasource `yaml:"datasource"`
}

type server struct {
	Env  string      `yaml:"env" validate:"required"`
	HTTP http.Config `yaml:"http"`
}

type datasource struct {
	UsedDB  string         `yaml:"used_db"`
	MySQL   mysql.Config   `yaml:"mysql"`
	MongoDB mongodb.Config `yaml:"mongodb"`
}
