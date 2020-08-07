// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"net/http"

	"github.com/alimy/alexandrite/internal/config"
	"github.com/alimy/alexandrite/mirc/gen/api"
	"github.com/alimy/alexandrite/servants"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	v1 "github.com/alimy/alexandrite/mirc/gen/api/api/v1"
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
	serveCmd.Flags().StringVarP(&inConfigFile, "config", "c", "", "custom config file path used to init application")
	register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	conf := serveSetup()
	r := mux.NewRouter()

	// register servants to chi
	registerServants(r)

	// start servant service
	if err := http.ListenAndServe(conf.Server.Addr, r); err != nil {
		logrus.Fatal(err)
	}
}

func registerServants(r *mux.Router) {
	api.RegisterFrontendServant(r, servants.NewFrontend())
	v1.RegisterRegistryServant(r, servants.NewRegistry())
}

func serveSetup() *config.Config {
	conf := config.InitWith(inConfigFile)
	coreInit(conf)
	return conf
}
