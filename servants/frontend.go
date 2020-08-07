package servants

import (
	"net/http"

	"github.com/alimy/alexandrite/internal/assets"
	"github.com/alimy/alexandrite/mirc/gen/api"
	"github.com/gorilla/mux"
)

type frontend struct {
	staticHandler http.Handler
	// TODO
}

func (f *frontend) Chain() []mux.MiddlewareFunc {
	return nil
}

func (f *frontend) Index(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Me(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func (f *frontend) Search(rw http.ResponseWriter, r *http.Request) {
	// TODO
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

func (f *frontend) Assets(rw http.ResponseWriter, r *http.Request) {
	// TODO
}

func NewFrontend() api.Frontend {
	fs := assets.NewFileSystem()
	return &frontend{
		staticHandler: http.StripPrefix("/assets/", http.FileServer(fs)),
	}
}
