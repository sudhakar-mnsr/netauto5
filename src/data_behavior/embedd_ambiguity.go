// Sample program to show what happens when the outer and
// inner type implement the same interface.
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior
type notifier interface {
   notify()
}

// user defines a user in the program
type user struct {
   name string
   email string
}

// user implements method of notifier interface
func (u *user) notify() {
   fmt.Printf("Sending user email to %s,%s>\n", u.name, u.email)
}

