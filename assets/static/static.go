// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package static

import (
	"embed"
	"net/http"
)

//go:embed fonts wasm icons.css
var resource embed.FS

// NewFS returns an http.FileSystem instance backed by embedded assets.
func NewFS() http.FileSystem {
	return http.FS(resource)
}
