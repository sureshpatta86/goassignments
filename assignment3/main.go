package main

import (
	"fmt"
	"go-assignments/assignment3/domain"
	"go-assignments/assignment3/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

// Add a new customer to the controller
func (c *CustomerController) Add(customer domain.Customer) error {
	err := c.store.CreateCustomer(customer)
	if err != nil {
		return err
	}
	fmt.Println("New Customer has been created")
	return nil
}

func main() {
	// Create a new controller
	controller := &CustomerController{
		store: mapstore.NewMapStore(),
	}

	// Create a new Customer
	customer := domain.Customer{
		ID:    "customer1",
		Name:  "Suresh",
		Email: "suresh@gmail.com",
	}

	// Add the customer to the controller
	err := controller.Add(customer)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
