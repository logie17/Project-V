package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"net/http"
)

var index_tmpl = pongo2.Must(pongo2.FromFile("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", "Logie Sucks")
}
