// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package store

import (
	"github.com/alimy/hori/dao"
	"github.com/dgraph-io/badger/v2"
	"github.com/sirupsen/logrus"
)

type bs struct {
	*badger.DB
}

func (s *bs) Whoami() string {
	return "badger"
}

func newBS(path string) dao.Stored {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		logrus.Fatal(err)
	}
	return &bs{
		DB: db,
	}
}
