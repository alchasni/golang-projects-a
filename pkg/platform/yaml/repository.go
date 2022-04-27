package yaml

import (
	"fmt"
	"io/ioutil"

	"twatter/pkg/core/adapter/validatoradapter"

	"gopkg.in/yaml.v3"
)

func Init(path string, validator validatoradapter.Adapter) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	valErr := validator.Struct(cfg)
	if valErr != nil {
		return nil, valErr
	}

	return cfg, nil
}
