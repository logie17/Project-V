package main

import (
	"github.com/flosch/pongo2"
	"net/http"
)

var index_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/index.html"))
var index_partial_tmpl = pongo2.Must(pongo2.FromFile("templates/partials/index.html"))

// TODO: test this sucker and figure out how to mock net/http
func indexHandler(w http.ResponseWriter, r *http.Request) {
	ajax := r.Header.Get("X-PUSH")
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
