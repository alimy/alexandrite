// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import (
	"github.com/alimy/hori/dao"
	"github.com/dgraph-io/ristretto"
	"github.com/sirupsen/logrus"
)

type mr struct {
	*ristretto.Cache
}

func (m *mr) Whoami() string {
	return "ristretto"
}

func newMC() dao.Cached {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 65535,   // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
		KeyToHash: func(key interface{}) (uint64, uint64) {
			return uint64(key.(uint16)), 0
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return &mr{
		Cache: cache,
	}
}

func (m *mr) PutConfig(content string) {
	m.Set(keyConfig, content, 1)
}

func (m *mr) GetConfig() (string, bool) {
	v, exist := m.Get(keyConfig)
	if !exist {
		return "", false
	}
	value, ok := v.(string)
	return value, ok
}
