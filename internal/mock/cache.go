// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mock

type cache struct{}

func (cache) Whoami() string {
	return "mockCached"
}

func (cache) PutConfig(string) {
	// Empty
}

func (cache) GetConfig() (string, bool) {
	return "", false
}
