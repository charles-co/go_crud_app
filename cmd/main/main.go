package main

import (
	_ "github.com/charles-co/go_crud_app/docs"
	"github.com/charles-co/go_crud_app/pkg/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title           Book Store API
// @version         1.0
// @description     This is a sample server for a book store.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http

func main() {
	r := mux.NewRouter()
	api_v1 := r.PathPrefix("/api/v1").Subrouter()
	routes.ResgisterBookStoreRoutes(api_v1)

	http.Handle("/", r)
	api_v1.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/api/v1/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
