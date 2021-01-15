// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"fmt"
	"io/fs"
	"path"
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

// Naming rename template name from filepath
func Naming(filepath string) string {
	ext := path.Ext(filepath)
	return filepath[:len(filepath)-len(ext)]
}
