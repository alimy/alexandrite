// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"

	"github.com/alimy/alexandrite/version"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Version of application",
		Run:   versionRun,
	}
	register(versionCmd)
}

func versionRun(cmd *cobra.Command, args []string) {
	verBuf := bytes.NewBufferString(fmt.Sprintf("v%s\n", version.AppVer))
	if version.BuildTime != "Not provided" && version.GitHash != "Not provided" {
		verBuf.WriteString(fmt.Sprintf("BuildTime:%s\nGitHash:%s\n", version.BuildTime, version.GitHash))
	}
	fmt.Printf("%s", verBuf)
}
