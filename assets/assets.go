// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"net/http"

	"github.com/alimy/embedx"
)

// DefaultConfig default configure raw data string.
func DefaultConfig() string {
	//go:embed conf/app.toml
	var content string

	return content
}

// NewStaticFS make a http.FileSystem instance that contain static files.
func NewStaticFS() http.FileSystem {
	//go:embed static
	var content embed.FS

	return http.FS(embedx.ChangeRoot(content, "static"))
}
