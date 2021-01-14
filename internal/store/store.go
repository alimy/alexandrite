// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package store

import (
	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
)

func NewStored(config *conf.Store) dao.Stored {
	switch config.Type {
	case "badger":
		return newBS(config.Path)
	default:
		return newBS(config.Path)
	}
}
