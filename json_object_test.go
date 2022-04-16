package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
- Json Object in golang as representation dari a `struct`, so struct a structure of Json in Golang
*/

// Struct for object Customer
type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

// Function test
func TestJsonObject(t *testing.T) {
	// (1) Added value struc
	customer := Customer{
		FirstName: "Rizky",
		LastName:  "Darmawan",
		Age:       28,
		Married:   true,
	}

	// (2) Create JSON
	bytes, err := json.Marshal(customer)

	// (3) If Error
	if err != nil {
		panic(err)
	}

	// (5) If no, print and convert from slice to string
	fmt.Println(string(bytes))
}
