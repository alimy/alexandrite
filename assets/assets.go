// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"net/http"

	"github.com/alimy/embedx"
)

var (
	//go:embed conf/app.toml
	configContent string

	//go:embed static
	staticFS embed.FS
)

// DefaultConfig default configure raw data string.
func DefaultConfig() string {
	return configContent
}

// NewStaticFS make a http.FileSystem instance that contain static files.
func NewStaticFS() http.FileSystem {
	return http.FS(embedx.ChangeRoot(staticFS, "static"))
}
