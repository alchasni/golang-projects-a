package yaml

import (
	"twatter/pkg/platform/mysql"
	"twatter/pkg/transport/http"
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
	MySQL mysql.Config `yaml:"mysql"`
}
