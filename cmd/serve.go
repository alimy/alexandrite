// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"net/http"

	"github.com/alimy/hori/internal/conf"
	"github.com/alimy/hori/servants"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	inConfigFile string
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start server service",
		Run:   serveRun,
	}
	serveCmd.Flags().BoolVar(&inDebug, "debug", false, "whether in debug mode")
	serveCmd.Flags().StringVarP(&inConfigFile, "config", "c", "custom/conf/app.toml", "custom config file path used to init application")
	register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	config := serveSetup()

	r := mux.NewRouter()
	servants.RegisterServants(r)

	logrus.Printf("start listening on %s", config.Server.Addr)
	if err := http.ListenAndServe(config.Server.Addr, r); err != nil {
		logrus.Fatal(err)
	}
}

func serveSetup() *conf.Config {
	config := conf.InitWith(inConfigFile)
	logrus.Debugf("config:%s", config)

	inSetup(config)
	coreInit(config)
	return config
}
