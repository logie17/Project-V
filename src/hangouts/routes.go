package main

import (
	"github.com/flosch/pongo2"
	"net/http"
)

var index_tmpl = pongo2.Must(pongo2.FromFile("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := index_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
