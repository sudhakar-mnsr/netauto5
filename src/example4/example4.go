package main

import (
   "fmt"
   "example4/emp"
)

func main() {
   m := emp.Manager{
      Teamsize: 10,
   }
   m.Name = "sudhakar"
   m.Sal = 2.0
   fmt.Println(m)
}
