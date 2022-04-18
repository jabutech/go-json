package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
- By default attribute on `Struct` and JSON will the mapping based on same attribute name (case sensitive)
- If in `struct` use snake_case or PascalCase then in JSON is same
- We can use `Tag Reflection` for then JSON name, and then followed by the name attribute
*/

// Struct with tag reflection
type Product struct {
	Id       string `json:"id"`        // Result of json = id
	Name     string `json:"name"`      // Result of json = name
	Price    int64  `json:"price"`     // Result of json = price
	ImageURL string `json:"image_url"` // Result of json = id
}

// Function Test Encode
func TestJSONTagEncode(t *testing.T) {
	// (1) Create new value for Object Product
	product := Product{
		Id:       "P0001",
		Name:     "Apple Macbook Pro",
		ImageURL: "http://example.com/image.png",
	}

	// (2) Convert to JSON
	bytes, err := json.Marshal(product)

	// (3) If error
	if err != nil {
		panic(err)
	}

	// (4) If no error, print and convert to string
	fmt.Println(string(bytes))
}

// Function test decode
func TestJSONTagDecode(t *testing.T) {
	// (1) Create data for decode
	jsonString := `{"id":"P0001","name":"Apple Macbook Pro","price":0,"image_url":"http://example.com/image.png"}`
	// (2) Convert to data type slice byte
	jsonBytes := []byte(jsonString)

	// (3) Get object struct product
	product := &Product{}

	// (4) Decode
	json.Unmarshal(jsonBytes, product)

	// (5) Print
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageURL)
}
