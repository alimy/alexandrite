// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"net/http"

	"github.com/alimy/hori/assets"
	"github.com/alimy/hori/internal/conf"
	"github.com/alimy/hori/mirc/auto/api"
	"github.com/alimy/hori/servants"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	v1 "github.com/alimy/hori/mirc/auto/api/api/v1"
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
	config := serveSetup()
	r := mux.NewRouter()
	registerServants(r)

	logrus.Printf("start listening on %s", config.Server.Addr)
	if err := http.ListenAndServe(config.Server.Addr, r); err != nil {
		logrus.Fatal(err)
	}
}

func registerServants(r *mux.Router) {
	api.RegisterFrontendServant(r, servants.NewFrontend())
	v1.RegisterRegistryServant(r, servants.NewRegistry())

	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assets.NewStaticFS()))
	r.PathPrefix("/assets/").Handler(assetsHandler)
}

func serveSetup() *conf.Config {
	config := conf.InitWith(inConfigFile)
	coreInit(config)
	inSetup()
	return config
}
