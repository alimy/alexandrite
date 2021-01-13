// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"sync"

	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
)

const (
	keyConfig uint16 = iota
)

var (
	cached dao.Cached
	once   sync.Once
)

func MyCached() dao.Cached {
	once.Do(func() {
		config := conf.MyConfig()
		switch config.Cache.Type {
		case "ristretto":
			cached = newMC()
		default:
			cached = newMC()
		}
	})
	return cached
}
