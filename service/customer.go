package service

import "github.com/pushm0v/gorest/model"

type CustomerService interface {
	GetCustomer(id int) (m *model.Customer)
}

type customerService struct {
	customers map [int]*model.Customer
}

func NewCustomerService() CustomerService {
	return &customerService{customers: map [int]*model.Customer{}}
}

func (c *customerService) GetCustomer(id int) (m *model.Customer) {
	return c.customers[id]
}


