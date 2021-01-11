// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cache

import "github.com/alimy/alexandrite/dao"

type mc struct {
	// TODO
}

func newMC() dao.Cached {
	return &mc{}
}
