package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	formatted := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "02/01/2006"
	dateStr := "25/11/2030"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)

	// add 1 more day in currentTime
	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println("new_date time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time: ", formatted_new_date)

}
