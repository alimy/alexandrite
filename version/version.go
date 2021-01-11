// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import "github.com/Masterminds/semver/v3"

var (
	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"

	AppVer = semver.MustParse("v0.1.0")
)
