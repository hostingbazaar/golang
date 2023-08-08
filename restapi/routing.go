package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmp).Methods("POST")
	r.HandleFunc("/employees", GetEmp).Methods("GET")
	r.HandleFunc("/getemp/{name}", GetEmpid).Methods("GET")
	r.HandleFunc("/updemp/{name}", UpdEmp).Methods("PUT")
	r.HandleFunc("/delete/{name}", DeleteEmp).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r)) //if you not use it will stop crawler

}
