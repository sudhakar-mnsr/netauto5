// Sample program to show how to embed a type into another
// type and the relationship between the inner and outer type.

package main

import "fmt"

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

// admin represents an amdin user with privileges
type admin struct {
   user // Embedded type
   level string
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
   
   // we can access the inner types method directly
   ad.user.notify()
   // The inner types method is promoted
   ad.notify()
}
