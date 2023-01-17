package main

import (
	"github.com/charles-co/go_crud_app/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.ResgisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
