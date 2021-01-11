// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"sync"

	"github.com/alimy/alexandrite/dao"
	"github.com/sirupsen/logrus"
)

var (
	repo dao.Repository
	once sync.Once
)

func MyRepo() dao.Repository {
	once.Do(func() {
		var err error
		repo, err = initDB()
		if err != nil {
			logrus.Fatal(err)
		}
	})
	return repo
}
