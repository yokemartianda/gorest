package main

import (
	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/service"
	"net/http"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()

	var custService = service.NewCustomerService()
	var custHandler = NewCustomerHandler(custService)
	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id}", custHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id}", custHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", custHandler.NotFound)
	r.Use(LoggingMiddleware)
	return r
}

