package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", indexHandler);
	http.ListenAndServe(":80", nil)
}

func indexHandler(r *http.Request, w http.ResponseWriter){
	// screw you vishal, (he doesnt want me to use a framework)
	fmt.Fprintf(w, "<h1>%s</h1>", "Logie Sucks")
}
