// Sample program to show how embedded types work with interfaces
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

// notify implements a method notifies users
// of different events
func (u *user) notify() {
   fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}


