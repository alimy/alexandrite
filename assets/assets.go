// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/alimy/embedx"
	"github.com/alimy/embedx/html"
)

// DefaultConfig default configure raw data string.
func DefaultConfig() string {
	//go:embed conf/hori.toml
	var content string

	return content
}

// NewStaticFS make a http.FileSystem instance that contain static files.
func NewStaticFS() http.FileSystem {
	//go:embed static
	var content embed.FS

	return http.FS(embedx.ChangeRoot(content, "static"))
}

// NewTemplate new template.Template instance from templates files.
func NewTemplate() (*template.Template, error) {
	//go:embed templates
	var content embed.FS

	embedFS := embedx.ChangeRoot(content, "templates")
	return html.ParseFS(embedFS, "*.tmpl", "partials/*.tmpl", "account/*.tmpl")
}
