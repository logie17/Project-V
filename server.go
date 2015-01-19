package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	"github.com/logie17/Project-V/config"
	"net/http"
	"log"
)

type key int
const UserKey key = 0

func main() {
	r := mux.NewRouter()
	configuration := config.LoadConfig()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", Middleware(r))

	fmt.Printf("Starting on Port: [::]:%v\n", configuration.Port)
	http.ListenAndServe(fmt.Sprintf("[::]:%s", configuration.Port), nil)
}

// Apparently gorrilla doesn't support some sort
// of route chaining or middleware :-(
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, UserKey, "foobar")
		log.Println("middleware begin")
		h.ServeHTTP(w, r)
		log.Println("middleware end")
	})
}
