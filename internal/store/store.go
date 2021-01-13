// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package store

import (
	"sync"

	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
)

var (
	stored dao.Stored
	once   sync.Once
)

func MyStored() dao.Stored {
	once.Do(func() {
		config := conf.MyConfig()
		switch config.Store.Type {
		case "badger":
			stored = newBS(config.Store.Path)
		default:
			stored = newBS(config.Store.Path)
		}
	})
	return stored
}
