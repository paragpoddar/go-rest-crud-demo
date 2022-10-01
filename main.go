package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello from go-rest-crud-demo app")
	}).Methods("GET")
	router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Person created")
	}).Methods("POST")
	router.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Got all Persons")
	}).Methods("GET")
	router.HandleFunc("/get/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Got Person by id")
	}).Methods("GET")
	router.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Updated Person by id")
	}).Methods("PUT")
	router.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Deleted Person by id")
	}).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8090", router))
}
