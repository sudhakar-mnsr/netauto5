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

// admin represents an admin user with privileges
type admin struct {
   user 
   level string
}

// notify implements a method notifies admins
// of different events.
func (a *admin) notify() {
   fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
// Create an admin user
ad := admin{
   user: user{
      name: "John Smith",
      email: "john@yahoo.com",
   },
   level: "super",
}

// Send the admin user a notification
// The embedded inner types implementation of the 
// interface is NOT "promoted" to the outer type.
sendNotification(&ad)
ad.user.notify()
ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
   n.notify()
} 
