// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
)

const (
	keyConfig uint16 = iota
)

func NewCached(config *conf.Cache) dao.Cached {
	switch config.Type {
	case "ristretto":
		return newMC()
	default:
		return newMC()
	}
}
