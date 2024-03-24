package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1, 30, 22, 19:
		fmt.Println("Good Day we are very fine")
	case 2:
		fmt.Println("Bad Day")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}

	temperature := 10
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
