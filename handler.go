package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/service"
	"log"
	"net/http"
	"strconv"
)

type CustomerHandler struct{
	custService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{custService: customerService}
}

func (s *CustomerHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	customer := s.custService.GetCustomer(custID)

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, customer)
}

func (s *CustomerHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "post called")
}

func (s *CustomerHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "put called")
}

func (s *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "delete called")
}

func (s *CustomerHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
