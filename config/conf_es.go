package config

import (
	"fmt"
)

type ES struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (e ES) URL() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
