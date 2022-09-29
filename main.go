package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// https://github.com/ariejanuarb/kabayan-backend-test

type Passanger struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// load the json file
	content, err := os.ReadFile("backend-titanic-test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// parse json file to []Passanger slice
	var passangers []Passanger
	err2 := json.Unmarshal(content, &passangers)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	fmt.Println(passangers)

	// find most frequent age numbers
	// print and sort apabethically the value of every most frequent numbers
}
