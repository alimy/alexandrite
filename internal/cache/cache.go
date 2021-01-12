// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"sync"

	"github.com/alimy/hori/dao"
)

var (
	cached dao.Cached
	once   sync.Once
)

func MyCached() dao.Cached {
	once.Do(func() {
		cached = newMC()
	})
	return cached
}
