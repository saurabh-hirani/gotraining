// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Add imports.

// Declare a type named user.

// User - user struct
type User struct {
	name        string
	email       string
	accessLevel int
}

// Create a function that changes the value of one of the user fields.
func setaccessLevel(u *User, accessLevel int) {
	// Use the pointer to change the value that the
	// pointer points to.
	u.accessLevel = accessLevel
}

func main() {

	// Create a variable of type user and initialize each field.
	var user1 User
	user1.name = "user1"
	user1.email = "user1@user1.com"
	user1.accessLevel = 0

	// Display the value of the variable.
	fmt.Println("access: ", user1.accessLevel)

	// Share the variable with the function you declared above.
	setaccessLevel(&user1, 12)

	// Display the value of the variable.
	fmt.Println("access: ", user1.accessLevel)
}
