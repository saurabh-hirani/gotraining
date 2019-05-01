// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup

	// Declare an anonymous function and create a goroutine.
	wg.Add(1)
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		defer wg.Done()
		for i := 100; i >= 0; i-- {
			fmt.Println("1: ", i)
		}
		// Tell main we are done.
		fmt.Println("100 to 1 completed")
	}()

	// Declare an anonymous function and create a goroutine.
	wg.Add(1)
	go func() {
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			fmt.Println("2: ", i)
		}
		// Tell main we are done.
		fmt.Println("1 to 100 completed")
	}()

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display "Terminating Program".
	fmt.Println("Terminating program")
}
