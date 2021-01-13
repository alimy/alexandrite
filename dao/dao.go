// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package dao

type Repository interface {
	Whoami() string
}

type Cached interface {
	Whoami() string
	PutConfig(string)
	GetConfig() (string, bool)
}

type Stored interface {
	Whoami() string
}

// Tables is the list of struct-to-table mappings.
func Tables() []interface{} {
	// TODO add struct to slice
	return []interface{}{}
}
