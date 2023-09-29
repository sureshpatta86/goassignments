package mapstore

import (
	"assignment3/domain"
	"errors"
	"sync"
)

type MapStore struct {
	mu    sync.RWMutex
	store map[string]domain.Customer
}

// NewMapStore will return a new instance
func NewMapStore() *MapStore {
	return &MapStore{
		store: make(map[string]domain.Customer),
	}
}

// Create Customer Interface
func (ms *MapStore) CreateCustomer(c domain.Customer) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.store[c.ID]; ok {
		return errors.New("customer already exist")
	}
	ms.store[c.ID] = c
	return nil
}

// Update Customer
func (ms *MapStore) UpdateCustomer(id string, c domain.Customer) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.store[id]; !ok {
		return errors.New("customer not found to update")
	}

	ms.store[id] = c
	return nil
}

// Delete Customer
func (ms *MapStore) DeleteCustomer(id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, ok := ms.store[id]; !ok {
		return errors.New("customer not found to delete")
	}

	delete(ms.store, id)
	return nil
}

// Get Cusotmer Details by ID
func (ms *MapStore) GetCustomerByID(id string) (domain.Customer, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	c, ok := ms.store[id]

	if !ok {
		return domain.Customer{}, errors.New("customer not found to get details")
	}

	return c, nil
}

// Get All Customers
func (ms *MapStore) GetAllCustomers() ([]domain.Customer, error) {

	ms.mu.RLock()
	defer ms.mu.RUnlock()

	customers := make([]domain.Customer, 0, len(ms.store))
	for _, c := range ms.store {
		customers = append(customers, c)
	}

	return customers, nil
}
