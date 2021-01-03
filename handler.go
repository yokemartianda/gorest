package main

import (
	"encoding/json"
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
	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "get called")
}

func (s *Handler) Post(w http.ResponseWriter, r *http.Request) {
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
