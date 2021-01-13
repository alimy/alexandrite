// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"sync"

	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
	"github.com/sirupsen/logrus"
)

var (
	repo dao.Repository
	once sync.Once
)

func MyRepo() dao.Repository {
	once.Do(func() {
		var err error
		config := conf.MyConfig()
		if config.Runtime.FakeDatabase {
			repo = fakeDB()
			return
		}
		repo, err = initDB()
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Infof("use %s(%s) as repository", repo.Whoami(), config.Database.Type)
	})
	return repo
}
