package main

import (
	"assignment3/domain"
	"fmt"
)

// CustomerController Organises the CRUD operations at UI layer.
type CustomerController struct {
	store domain.CustomerStore
}

// Add function to add new customer.
func (cc *CustomerController) Add(c domain.Customer) error {
	err := cc.store.CreateCustomer(c)
	if err != nil {
		return err
	}
	fmt.Println("New Customer has been created")
	return nil
}

// Update function to update a record.
func (cc *CustomerController) Update(s string, c domain.Customer) error {
	err := cc.store.UpdateCustomer(s, c)
	if err != nil {
		return err

	}
	fmt.Printf("Customer %s record has been updated\n", s)
	return nil
}

// Remove function to remove a record.
func (cc CustomerController) Remove(s string) error {
	err := cc.store.DeleteCustomer(s)
	if err != nil {
		return err
	}
	fmt.Printf("Customer %s record has been deleted\n", s)
	return nil
}

// GetByCustomerId to get individual record based on id.
func (cc CustomerController) GetByID(s string) error {
	c, err := cc.store.GetCustomerByID(s)
	if err != nil {
		return err
	}
	fmt.Printf("ID: %s\nName: %s\nE-mail:%s\n", c.ID, c.Name, c.Email)
	return nil
}

// GetAll to get all customer records.
func (cc CustomerController) GetAll()  error {
	cs, err := cc.store.GetAllCustomers()
	if err != nil {
		return  err
	}
	for _, c := range cs {
		fmt.Printf("ID: %s\nName: %s\nE-mail:%s\n", c.ID, c.Name, c.Email)
	}
	return nil
}
