package http

type Config struct {
	Port                string `yaml:"port" validator:"required"`
	Debug               bool   `yaml:"debug"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds" validator:"required"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds" validator:"required"`
	BaseURL             string `yaml:"base_url" validator:"required"`
}
