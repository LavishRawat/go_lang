package main

import (
	"log"
	"net/http"
	routes "Server/src/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)


func main() {
	r := mux.NewRouter()
	routes.RegisterSignUpDetailsRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}