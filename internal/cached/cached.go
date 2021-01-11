// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cached

import (
	"sync"

	"github.com/alimy/alexandrite/dao"
)

var (
	cache dao.Cache
	once  sync.Once
)

func MyCache() dao.Cache {
	once.Do(func() {
		cache = newMC()
	})
	return cache
}
