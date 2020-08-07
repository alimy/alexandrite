// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	v1 "github.com/alimy/alexandrite/mirc/gen/api/api/v1"
)

type registry struct {
	// TODO
}

func (g *registry) Register(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) Login(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) TokensInfo(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) GenerateTokens(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) RevokeTokens(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) TokenByName(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) Categories(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) SearchCrates(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) PublishCrates(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) SuggestCrates(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) CrateInfo(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) CrateOwnersInfo(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) PutCrateOwners(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) DelCrateOwners(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) YankCrate(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) UnyankCrate(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (g *registry) DownloadCrate(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func NewRegistry() v1.Registry {
	return &registry{}
}
