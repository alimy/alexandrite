// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/alexandrite/internal/config"
	"github.com/alimy/alexandrite/internal/logus"
)

var (
	inDebug bool
)

func inSetup() {
	if inDebug {
		logus.SetLevel(logus.LevelDebug)
	} else {
		logus.SetLevel(logus.LevelInfo)
	}
}

func coreInit(conf *config.Config) {
	// TODO
}
