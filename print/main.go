package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	//age: 25height: 5.8234567name: Alice
	fmt.Println("Hello world")

	// fmt.Printf("age is %d\n", age)
	// fmt.Printf("height is %.2f\n", height)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
