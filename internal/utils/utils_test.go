// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import "testing"

func TestNaming(t *testing.T) {
	for origin, expect := range map[string]string{
		"templates/a.tmpl":   "templates/a",
		"templates/b/c.tmpl": "templates/b/c",
		"templates/d/e.tmpl": "templates/d/e",
		"templates/d/f.tmpl": "templates/d/f",
	} {
		if name := Naming(origin); name != expect {
			t.Errorf("expect: %s got %s", expect, name)
		}
	}
}
