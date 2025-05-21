package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var err error

func main() {
	//create  a router
	r := mux.NewRouter()

	r.HandleFunc("/healtz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "App is Up and Running Successfully")
	}).Methods("Get")

	if err = http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
