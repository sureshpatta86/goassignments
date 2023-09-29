package domain

//Declare Customer Struct

type Customer struct {
	ID    string
	Name  string
	Email string
}

//Declare Customer Interface

type CustomerStore interface {
	CreateCustomer(Customer) error
	UpdateCustomer(string, Customer) error
	DeleteCustomer(string) error
	GetCustomerByID(string) (Customer, error)
	GetAllCustomers() ([]Customer, error)
}
