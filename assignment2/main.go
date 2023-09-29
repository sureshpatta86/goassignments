package main

import (
	"fmt"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

// addItem
func addItem(k, v string) {
	data[k] = v
}

// updateItem
func updateItem(k, v string) {
	if _, ok := data[k]; ok {
		data[k] = v
	} else {
		fmt.Println("Key not found")
	}
}

// getById
func getById(k string) string {
	return data[k]
}

// getAll
func getAll() map[string]string {
	return data
}

// deleteItem
func deleteItem(k string) {
	delete(data, k)
}

func main() {
	// Test the functions
	addItem("key1", "value1")
	addItem("key2", "value2")

	fmt.Println("Initial Map:", getAll())

	updateItem("key1", "newvalue1")
	fmt.Println("Updated Map:", getAll())

	fmt.Println("Value for key 'key2':", getById("key2"))

	deleteItem("key2")
	fmt.Println("Map after deleting 'key2':", getAll())

}
