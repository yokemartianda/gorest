package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"log"
	"net/http"
)

type Handler struct{}

func (s *Handler) responseBuilder(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for k,v := range vars {
		fmt.Printf("Key : %v, Value : %v\n", k, v)
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "get called")
}

func (s *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Address string `json:"address"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatalf("Request decoder error : %v", err)
	}

	fmt.Printf("%+v", req)

	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "post called")
}

func (s *Handler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "put called")
}

func (s *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "delete called")
}

func (s *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
