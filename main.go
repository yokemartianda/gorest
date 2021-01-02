package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func (s *Server) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	var s = &Server{}
	log.Println("Running http server")
	r := mux.NewRouter()
	r.HandleFunc("/", s.Get).Methods(http.MethodGet)
	r.HandleFunc("/", s.Post).Methods(http.MethodPost)
	r.HandleFunc("/", s.Put).Methods(http.MethodPut)
	r.HandleFunc("/", s.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", s.NotFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}