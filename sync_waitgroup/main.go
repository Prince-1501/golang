package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("worker %d started\n", i)
	// some task is happening
	fmt.Printf("worker %d end\n", i)
}

// worker(1)
// worker(2)
// worker(3)

func main() {
	// fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("workers task complete")

}
