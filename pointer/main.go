package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	// var num int
	num := "Prince"

	// var ptr *int
	// ptr = &num

	ptr := &num

	// fmt.Println("Num has value: ", num)
	fmt.Println("pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer: ", *ptr)

	var pointer *int // default pointer == nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 5
	modifyValueByReference(&value)
	fmt.Println("Value contains : ", value)

}
