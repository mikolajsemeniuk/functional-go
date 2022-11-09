package cart

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Listen string `envconfig:"LISTEN" default:":5000"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}

	return config, nil
}
