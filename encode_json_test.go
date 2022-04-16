package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
- For data conversion to JSON use a function `json.Marshal()`
- Data returned as interface{}, so we can send any data type
- There are two variable returned, `bytes` and `err`
*/

func logJson(data interface{}) {
	// (1) Convert parameter data to json
	bytes, err := json.Marshal(data)
	// (2) If error
	if err != nil {
		// returned panic
		panic(err)
	}
	// (3) If no error, print json data, and convertion byte from slice to string
	fmt.Println(string(bytes))
}

// Function testing
func TestEndcode(t *testing.T) {
	// (1) Send argument to func LogJson
	logJson("Hai")
	logJson(28)
	logJson(true)
	logJson([]string{"Rizky", "Darmawan"})
}
