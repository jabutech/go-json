package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
- Decode is convert data from JSON to data type of Golang
- To decode can use function `json.Unmarshal(byte[], interface{})`
- `byte[]` as data JSON which will be decoded to Golang
- `interface{}` as a place to store the results of the konversi, that is a pointer

*/

func TestDecodeJSON(t *testing.T) {
	// (1) Create a data JSON
	jsonString := `{"FirstName":"Rizky","LastName":"Darmawan","Age":3,"Married":true}`
	// (2) Convert data to byte
	jsonBytes := []byte(jsonString)

	// (3) Create an Empty struct and set a pointer
	customer := &Customer{}

	// (4) Decoded
	err := json.Unmarshal(jsonBytes, customer)

	// (5) If Error
	if err != nil {
		panic(err)
	}

	// (6) If no print
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
