// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

//go:generate go-bindata -nomemcopy -pkg=config -ignore=README.md -prefix=../../assets/ -debug=false -o=config_gen.go ../../assets/config/...

var (
	config *Config
)

// InitWith initialize models with custom config file
func InitWith(path string) *Config {
	config = &Config{}
	// init config
	if err := loadConfig(config); err != nil {
		logrus.Error("load config error", err)
	}
	if path == "" {
		// Empty
	} else if fileIsExist(path) { // set config from custom config file
		customConfig(config, path)
	} else {
		logrus.Infof("custom config file is not exist so use default config path: %s", path)
	}
	return config
}

// MyConfig return application's config
func MyConfig() *Config {
	return config
}

// fileIsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func loadConfig(config *Config) error {
	data := MustAsset("config/alexandrite.toml")
	_, err := toml.Decode(string(data), config)
	return err
}

func customConfig(config *Config, path string) {
	customConfig := &Config{}
	meta, err := toml.DecodeFile(path, customConfig)
	if err != nil {
		logrus.Errorf("decode custom config error %v", err)
		return
	}
	for _, key := range meta.Keys() {
		if len(key) == 1 { // top section just continue
			continue
		}
		switch key[0] {
		case "application":
			myAppSection(config, customConfig, key)
		case "server":
			myServerSection(config, customConfig, key)
		}
	}
}

func myAppSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "name":
		config.Application.Name = custom.Application.Name
	case "authors":
		config.Application.Authors = custom.Application.Authors
	case "description":
		config.Application.Description = custom.Application.Description
	}
}

func myServerSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "addr":
		config.Server.Addr = custom.Server.Addr
	}
}
