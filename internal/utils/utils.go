// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"fmt"
	"io/fs"
	"path"
	"strings"
)

// FileGlob files name in fs.FS that match patterns
func FileGlob(fsys fs.FS, patterns ...string) ([]string, error) {
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		filenames = append(filenames, list...)
	}
	return filenames, nil
}

// PartialName returns full base file name without file ext
// example: /foo/bar/baz.png => foo/bar/baz
func PartialName(filePath string) string {
	fileExt := path.Ext(filePath)
	return strings.TrimLeft(filePath[:len(filePath)-len(fileExt)], "/")
}
