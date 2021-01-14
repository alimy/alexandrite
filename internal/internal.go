// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/cache"
	"github.com/alimy/hori/internal/conf"
	"github.com/alimy/hori/internal/mock"
	"github.com/alimy/hori/internal/store"
	"github.com/alimy/hori/internal/xorm"
	"github.com/sirupsen/logrus"
)

var (
	repo   dao.Repository
	stored dao.Stored
	cached dao.Cached
)

func InitWith(config *conf.Config) {
	if !config.Runtime.MockDatabase {
		repo = xorm.NewRepo(&config.Database)
	} else {
		repo = mock.NewRepo()
	}
	logrus.Debugf("use %s as Repository instance", repo.Whoami())

	if !config.Runtime.MockStore {
		stored = store.NewStored(&config.Store)
	} else {
		stored = mock.NewStored()
	}
	logrus.Debugf("use %s as Store instance", stored.Whoami())

	if !config.Runtime.MockCache {
		cached = cache.NewCached(&config.Cache)
	} else {
		cached = mock.NewCached()
	}
	logrus.Debugf("use %s as Cache instance", cached.Whoami())
}

func MyRepo() dao.Repository {
	return repo
}

func MyStored() dao.Stored {
	return stored
}

func MyCached() dao.Cached {
	return cached
}
