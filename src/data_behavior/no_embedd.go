// Sample program to show how/what we are doing is NOT 
// Embedding a type but just using a type as field

package main

import "fmt"

// user defines a user in the program.
type user struct {
   name string
   email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
   fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}
