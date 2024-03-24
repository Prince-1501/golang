package main

import (
	"fmt"
)

func main() {

	// numbers := []int{1, 2}
	// numbers = append(numbers, 3, 10, 12, 13, 14, 15, 16, 17, 18, 19)
	// fmt.Println("Number : ", numbers)
	// fmt.Printf("Number has data type : %T\n", numbers)
	// fmt.Println("Length : ", len(numbers))

	// name := []string{}
	// fmt.Println("name : ", name)

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))


	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
	
}
