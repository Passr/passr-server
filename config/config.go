package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Config struct {
	Port           string
	DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Server   string
	Port     string
	Username string
	Password string
}

func Load(path string) (config Config, err error) {
	tomlData, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if _, err := toml.Decode(string(tomlData), &config); err != nil {
		return config, err
	}

	return
}
