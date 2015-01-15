package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// screw you vishal, (he doesnt want me to use a framework)
	fmt.Fprintf(w, "<h1>%s</h1>", "Logie Sucks")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)
}
