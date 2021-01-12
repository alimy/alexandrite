// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/hori/cmd"
)

func main() {
	cmd.Setup(
		"hori",
		"rust crates registry",
		"rust crates registry",
	)
	cmd.Execute()
}
