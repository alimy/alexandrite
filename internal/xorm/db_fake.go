// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import "github.com/alimy/hori/dao"

type dbf struct{}

func fakeDB() dao.Repository {
	return &dbf{}
}

func (dbf) Whoami() string {
	return "fakeDB"
}
