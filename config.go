package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Layouts map[string]layout `toml:"layout"`
}

type layout struct {
	Length []int `toml:"length"`
}

func loadConfigToml() (tomlConfig, error) {

	var cfg tomlConfig

	tomlPath := filepath.Join(os.Getenv("HOME"), ".config", "cvfv", "config.toml")
	// settings read
	_, err := toml.DecodeFile(tomlPath, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("load config error :%s", err)
	}

	return cfg, nil

}
