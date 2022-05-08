// just me playing around with Json
// trying to recreate the syntax from https://en.wikipedia.org/wiki/JSON#Syntax

package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	IsAlive   bool   `json:"Is Alive"`
	Age       int    `json:"Age"`
	Address   address
}

type address struct {
	Street string `json:"Street"`
	City   string `json:"City"`
	ZIP    int    `json:"ZIP Code"`
}

func main() {

	employees := []employee{
		{FirstName: "John", IsAlive: true, Address: address{Street: "Mullholland Drive"}},
		{LastName: "Doe", IsAlive: false, Age: 23, Address: address{City: "NYC"}},
	}

	for i := range employees {
		bytes, _ := json.MarshalIndent(employees[i], "", "  ")
		fmt.Println(string(bytes))
	}
}
