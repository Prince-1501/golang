package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	// time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Prince :)")
	time.Sleep(1000 * time.Millisecond) // Simulating some work
	fmt.Println("Hi Prince Function ended:)")
}

func main() {
	fmt.Println("learning goroutines")

	go sayHello()
	go sayHi()

	// Wait for a moment to allow the goroutine to finish
	time.Sleep(800 * time.Millisecond)
}
