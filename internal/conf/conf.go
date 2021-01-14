// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/alimy/hori/assets"
	"github.com/sirupsen/logrus"
)

var (
	config = mustConfig()

	// Indicates which database backend is currently being used.
	UseSQLite3    bool
	UseMySQL      bool
	UsePostgreSQL bool
	UseMSSQL      bool
)

// InitWith initialize models with custom config file
func InitWith(path string) *Config {
	// set config from custom config file
	if path != "" && fileIsExist(path) {
		customConfig(config, path)
	} else {
		logrus.Infof("custom config file[%s] is not exist so use default config", path)
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

func mustConfig() *Config {
	defConf := &Config{}
	if _, err := toml.Decode(assets.DefaultConfig(), defConf); err != nil {
		panic(err)
	}
	return defConf
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
		case "runtime":
			myRuntimeSection(config, customConfig, key)
		case "server":
			myServerSection(config, customConfig, key)
		case "database":
			myDatabaseSection(config, customConfig, key)
		case "cache":
			myCacheSection(config, customConfig, key)
		case "store":
			myStoreSection(config, customConfig, key)
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

func myRuntimeSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "run_mode":
		config.Runtime.RunMode = custom.Runtime.RunMode
	case "mock_database":
		config.Runtime.MockDatabase = custom.Runtime.MockDatabase
	case "mock_store":
		config.Runtime.MockStore = custom.Runtime.MockStore
	case "mock_cache":
		config.Runtime.MockCache = custom.Runtime.MockCache
	}
}

func myServerSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "addr":
		config.Server.Addr = custom.Server.Addr
	}
}

func myDatabaseSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "type":
		config.Database.Type = custom.Database.Type
	case "host":
		config.Database.Host = custom.Database.Host
	case "name":
		config.Database.Name = custom.Database.Name
	case "user":
		config.Database.User = custom.Database.User
	case "password":
		config.Database.Password = custom.Database.Password
	case "path":
		config.Database.Path = custom.Database.Path
	case "ssl_mode":
		config.Database.SSLMode = custom.Database.SSLMode
	case "max_open_conns":
		config.Database.MaxOpenConns = custom.Database.MaxOpenConns
	case "max_idle_conns":
		config.Database.MaxIdleConns = custom.Database.MaxIdleConns
	}
}

func myCacheSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "type":
		config.Cache.Type = custom.Cache.Type
	}
}

func myStoreSection(config *Config, custom *Config, key toml.Key) {
	switch key[1] {
	case "type":
		config.Store.Type = custom.Store.Type
	case "path":
		config.Store.Path = custom.Store.Path
	}
}

// IsProdMode returns true if the application is running in production mode.
func IsProdMode() bool {
	return strings.EqualFold(config.Runtime.RunMode, "prod")
}
