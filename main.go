// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/alimy/alexandrite/mirc/gen/api"
	"github.com/alimy/alexandrite/servants"
	"github.com/gorilla/mux"

	v1 "github.com/alimy/alexandrite/mirc/gen/api/api/v1"
)

func main() {
	r := mux.NewRouter()

	// register servants to chi
	registerServants(r)

	// start servant service
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func registerServants(r *mux.Router) {
	api.RegisterFrontendServant(r, servants.NewFrontend())
	v1.RegisterRegistryServant(r, servants.NewRegistry())
}
