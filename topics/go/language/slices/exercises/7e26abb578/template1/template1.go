// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	// Display each value in the slice.
	for i := 0; i < 10; i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("*************")

	// Declare a slice of strings and populate the slice with names.
	var names []string
	names = append(names, "user1")
	names = append(names, "user2")
	names = append(names, "user3")

	// Display each index position and slice value.
	for idx, val := range names {
		fmt.Printf("names[%d] = %s\n", idx, val)
	}
	fmt.Println("*************")

	// Take a slice of index 1 and 2 of the slice of strings.
	firsttwo := names[1:3]

	// Display each index position and slice values for the new slice.
	for idx, val := range firsttwo {
		fmt.Printf("names[%d] = %s\n", idx, val)
	}

}
