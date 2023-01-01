package config

import "strings"

type Config struct {
	TLS bool
}

func NewConfig(args []string) (*Config, error) {
	if len(args) == 1 {
		return &Config{}, nil
	}

	var config Config
	for _, arg := range args {
		if strings.EqualFold(arg, "tls") {
			config.TLS = true
		}
	}

	return &config, nil
}
