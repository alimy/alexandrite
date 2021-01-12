// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/hori/internal/conf"
	"github.com/alimy/hori/internal/logus"
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

func coreInit(conf *conf.Config) {
	// TODO
}
