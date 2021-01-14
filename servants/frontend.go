// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/hori/assets"
	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal"
	"github.com/alimy/hori/mirc/auto/api"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type frontend struct {
	tmpl   assets.Template
	repo   dao.Repository
	cached dao.Cached
}

func (f *frontend) Chain() []mux.MiddlewareFunc {
	return nil
}

func (f *frontend) Index(rw http.ResponseWriter, r *http.Request) {
	// TODO
	logrus.Infof("get index %s", r.URL.Path)
	rw.Write([]byte(r.URL.Path))
}

func (f *frontend) Me(rw http.ResponseWriter, r *http.Request) {
	// TODO
	logrus.Infof("get me %s", r.URL.Path)
	rw.Write([]byte(r.URL.Path))
}

func (f *frontend) Search(rw http.ResponseWriter, r *http.Request) {
	// TODO
	logrus.Infof("get search %s", r.URL.RawPath)
	rw.Write([]byte(r.URL.Path))
}

func (f *frontend) MostDownloaded(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) LastUpdated(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Crate(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Login(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) LoginPost(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Logout(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Register(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) RegisterPost(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Manage(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Password(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Tokens(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) RevokeToken(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func NewFrontend() api.Frontend {
	return &frontend{
		tmpl:   assets.NewTemplate(),
		repo:   internal.MyRepo(),
		cached: internal.MyCached(),
	}
}
