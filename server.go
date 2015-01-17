package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kmulvey/Project-V/config"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	configuration := config.LoadConfig()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)
	fmt.Printf("Starting on Port: :%v\n", configuration.Port)
	fmt.Printf("Starting on Port: [::]:%v\n", configuration.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", configuration.Port), nil)
	http.ListenAndServe(fmt.Sprintf("[::]:%s", configuration.Port), nil)
}
