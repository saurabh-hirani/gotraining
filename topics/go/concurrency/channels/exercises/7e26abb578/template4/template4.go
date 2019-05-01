// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	maxCount := 100

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	// shutdown := make(chan string, 1)
	shutdown := make(chan struct{}) // only used for closing hence 0 size struct

	// Define the size of the worker pool. Use runtime.NumCPU to size the pool based on number of processors.
	workerCount := runtime.NumCPU()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < workerCount; i++ {

		// Start an infinite loop.
		go func(id int) {
			fmt.Printf("Worker %d - starting\n", id)
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {
				// In one case send the random number.
				case values <- n:
					fmt.Printf("Worker %d - sending %d\n", id, n)

				// In another case receive from the shutdown channel.
				case <-shutdown:
					wg.Done()
					fmt.Printf("Worker %d - stopping\n", id)
					return

				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	var nums []int

	// Receive from the values channel with range.
	for v := range values {
		// continue the loop if the value was even.
		if v%2 == 0 {
			continue
		}
		// Store the odd number.
		nums = append(nums, v)
		// break the loop once we have 100 results.
		if len(nums) == maxCount {
			break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	for _, val := range nums {
		fmt.Println(val)
	}
}
