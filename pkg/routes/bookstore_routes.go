package routes

import (
	"github.com/charles-co/go_crud_app/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

var ResgisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
