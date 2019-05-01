// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

func main() {

	// Declare and make a map of integer type values.
	ages := make(map[string]int)

	// Initialize some data into the map.
	ages["u1"] = 12
	ages["u2"] = 13

	// Display each key/value pair.
	ages["u4"]++
	for key, value := range ages {
		fmt.Printf("ages[%s] = %d\n", key, value)
	}
	fmt.Printf("ages[u5] = %d\n", ages["u5"])

	val, ok := ages["u6"]
	if ok {
		fmt.Printf("ages[u6] = %d\n", val)
	} else {
		fmt.Println("u6 not present")
	}
}
