package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a int, b int) (result int) {
	result = a + b
	return
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of two number is :", ans)

	data := multiply(3, 4)
	fmt.Println("multiplication of two numbers is :", data)
}
