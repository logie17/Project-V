package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	// screw you vishal, (he doesnt want me to use a framework)
	fmt.Fprintf(w, "<h1>%s</h1>", "Logie Sucks")
}
func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":1234", nil)
}

