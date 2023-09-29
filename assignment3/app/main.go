package main

import (
	"assignment3/domain"
	"assignment3/mapstore"
	"fmt"
)

func main() {
	// Create a new controller
	controller := CustomerController{ // initialize customer controller
		store: mapstore.NewMapStore(),
	}

	// Create a new Customer
	customer1 := domain.Customer{
		ID:    "customer1",
		Name:  "Suresh",
		Email: "suresh@gmail.com",
	}

	customer2 := domain.Customer{
		ID:    "customer2",
		Name:  "Nani",
		Email: "nani@gmail.com",
	}

	customer3 := domain.Customer{
		ID:    "customer3",
		Name:  "Rajesh",
		Email: "",
	}

	err := controller.Add(customer1) // Create a new Customer
	if err != nil {
		fmt.Println("Error:", err)
	}

	err1 := controller.Add(customer2) // Create a new Customer
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	err2 := controller.Add(customer3) // Create a new Customer
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	// Get Customer by ID
	err3 := controller.GetByID("customer1")
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

	// Get All Customers
	err4 := controller.GetAll()
	if err4 != nil {
		fmt.Println("Error:", err4)
	}

	// Update Customer
	customer4 := domain.Customer{
		ID:    "customer1",
		Name:  "Suresh",
		Email: "sureshnew@gmail.com",
	}
	controller.Update("customer1", customer4)

	// Get Customer by ID
	err5 := controller.GetByID("customer1")
	if err5 != nil {
		fmt.Println("Error:", err5)
	}

	// Delete Customer
	controller.Remove("customer3")

	// Get All Customers
	err6 := controller.GetAll()
	if err6 != nil {
		fmt.Println("Error:", err6)
	}

}
