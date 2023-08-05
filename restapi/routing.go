package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmp).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r)) //if you not use it will stop crawler

}
