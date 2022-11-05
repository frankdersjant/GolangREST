package main

import (
	"example/rest/Domain"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	customer := Domain.Customer{ID: "1", Name: "John", Address: "123 Main St", DateofBirth: "01/01/2000", Status: "Active"}

	fmt.Println(customer)
}
