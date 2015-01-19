package main

import (
	"github.com/flosch/pongo2"
	"net/http"
	"github.com/gorilla/context"
	"log"
)

var index_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/index.html"))
var index_partial_tmpl = pongo2.Must(pongo2.FromFile("templates/partials/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ajax := r.Header.Get("X-PUSH")
	log.Println(context.Get(r, UserKey))
	if ajax != "" {
		// serve the partial
		err := index_partial_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// serve a page
		err := index_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginTmpl := pongo2.Must(pongo2.FromFile("templates/pages/login.html"))
	err := loginTmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
