// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/hori/assets"
	"github.com/alimy/hori/mirc/auto/api"
	"github.com/gorilla/mux"

	v1 "github.com/alimy/hori/mirc/auto/api/api/v1"
)

func RegisterServants(r *mux.Router) {
	api.RegisterFrontendServant(r, newFrontend())
	v1.RegisterRegistryServant(r, newRegistry())

	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assets.NewStaticFS()))
	r.PathPrefix("/assets/").Handler(assetsHandler)
}
