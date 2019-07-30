package main

import (
	"log"
	"net/http"

	"./controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", controller.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
