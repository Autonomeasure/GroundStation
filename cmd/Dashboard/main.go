package main

import (
	v0 "github.com/Autonomeasure/GroundStation/cmd/Dashboard/httpHandlers/v0"
	"github.com/Autonomeasure/GroundStation/cmd/Dashboard/middleware"
	"github.com/Autonomeasure/GroundStation/pkg/Database"
	"github.com/Autonomeasure/GroundStation/pkg/Memory"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	setupRouter(true)
}

func setupRouter(debug bool) {
	Memory.Database = Database.Database{}
	Memory.Database.Open()
	// Create the router
	router := mux.NewRouter().StrictSlash(true)

	// Create the /api route
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Set all the middleware for the /api route
	apiRouter.Use(middleware.SetResponseTypeToJSON)
	apiRouter.Use(middleware.EnableCORS)
	if debug == true {
		apiRouter.Use(middleware.Logger)
	}

	// Create subrouters (api versions)
	apiV0Router := apiRouter.PathPrefix("/v0").Subrouter()

	// Add routes to /api/v0/
	apiV0Router.HandleFunc("", v0.Api)

	log.Fatal(http.ListenAndServe(":8080", router))
}