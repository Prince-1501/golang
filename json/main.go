package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("we are learning JSON")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	// fmt.Println("person Data is : ", person)

	// convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling ", err)
		return
	}
	fmt.Println("person Data is : ", string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling ", err)
		return
	}
	fmt.Println("person Data is after unmarshalling: ", personData)

}
