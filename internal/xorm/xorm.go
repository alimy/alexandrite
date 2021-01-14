// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
	"github.com/sirupsen/logrus"
)

func NewRepo(config *conf.Database) dao.Repository {
	repo, err := initDB()
	if err != nil {
		logrus.Fatal(err)
	}
	return repo
}
