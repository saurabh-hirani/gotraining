// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.

// User - user structure
type User struct {
	name string
	age  int
}

func main() {
	// Declare variable of type user and init using a struct literal.
	u1 := User{
		name: "user1",
		age:  12,
	}

	// Display the field values.
	fmt.Printf("%+v\n", u1)

	// Declare a variable using an anonymous struct.
	u1 = struct {
		name string
		age  int
	}{
		name: "user2",
		age:  13,
	}

	// Display the field values.
	fmt.Printf("%+v\n", u1)
}
