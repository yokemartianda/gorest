package service

import (
	"fmt"
	"github.com/pushm0v/gorest/model"
)

type CustomerService interface {
	GetCustomer(id int) (m *model.Customer)
	CreateCustomer(cust *model.Customer)
	UpdateCustomer(id int, cust *model.Customer) error
	DeleteCustomer(id int) error
}

type customerService struct {
	customers map [int]*model.Customer
}

func NewCustomerService() CustomerService {
	return &customerService{customers: map [int]*model.Customer{}}
}

func (c *customerService) generateCustID() int {
	lenOfCustomer := len(c.customers)
	return lenOfCustomer + 1
}

func (c *customerService) checkCustomerID(id int) error {
	if c.customers[id] == nil {
		return fmt.Errorf("Customer not found with ID %v", id)
	}

	return nil
}

func (c *customerService) GetCustomer(id int) (m *model.Customer) {
	return c.customers[id]
}

func (c *customerService) CreateCustomer(cust *model.Customer) {
	id := c.generateCustID()
	cust.ID = id

	c.customers[id] = cust
}

func (c *customerService) UpdateCustomer(id int, cust *model.Customer) error {
	err := c.checkCustomerID(id)
	if err != nil {
		return err
	}

	cust.ID = id
	c.customers[id] = cust
	return nil
}

func (c *customerService) DeleteCustomer(id int) error {
	err := c.checkCustomerID(id)
	if err != nil {
		return err
	}

	delete(c.customers,id)
	return nil
}


