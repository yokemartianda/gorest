package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	var s = &Handler{}
	r.HandleFunc("/", s.Get).Methods(http.MethodGet)
	r.HandleFunc("/", s.Post).Methods(http.MethodPost)
	r.HandleFunc("/", s.Put).Methods(http.MethodPut)
	r.HandleFunc("/", s.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", s.NotFound)
	r.Use(LoggingMiddleware)
	return r
}

