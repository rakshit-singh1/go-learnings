package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("worker %d started\n", i)
	// some task is happening
	time.Sleep(5 * time.Second)
	fmt.Printf("worker %d end\n", i)
}

// func worker(i int) {

// 	fmt.Printf("worker %d started\n", i)
// 	// some task is happening
// 	time.Sleep(5 * time.Second)
// 	fmt.Printf("worker %d end\n", i)
// }

func main() {
	// fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, &wg)
	wg.Add(1)
	go worker(2, &wg)
	wg.Add(1)
	go worker(3, &wg)

	// go worker(1)
	// go worker(2)
	// go worker(3)

	// // Start 3 worker goroutines
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1) // Increment the WaitGroup counter
	// 	go worker(i, &wg)
	// }

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("workers task complete")

}
