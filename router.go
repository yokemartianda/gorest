package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yokemartianda/gorest/model"
	"github.com/yokemartianda/gorest/repository"
	"github.com/yokemartianda/gorest/service"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	customerRouter(api)
	r.Use(LoggingMiddleware)
	return r
}

func customerRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection("customer.db")
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Customer{})
	var custRepository = repository.NewCustomerRepository(dbConn)
	var custService = service.NewCustomerService(custRepository)
	var custHandler = NewCustomerHandler(custService)

	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id}", custHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id}", custHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/customers", custHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.GetAll).Queries("limit", "{limit}", "offset", "{offset}").Methods(http.MethodGet)
	r.HandleFunc("/", custHandler.NotFound)
}
