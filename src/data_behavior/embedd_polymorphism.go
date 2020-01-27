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

// admin represents an admin user with privileges
type admin struct {
   user
   level string
}

func main() {
// Create admin user
ad := admin{
   user: user{
      name: "John Smith",
      email: "john@yahoo.com",
   },
   level: "super",
}

// Send the admin user a notification.
// The embedded inner type implementation of the 
// interface is "promoted" to the outer type.
sendNotification(&ad)
}

// SendNotification accepts values that implement the notifier
// interface and sends notifications
func sendNotification(n notifier) {
   n.notify()
}
