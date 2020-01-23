package main

import (
	"fmt"
)

/*
class Animal
   virtual abstract Speak() string
*/
type Animal interface {
	Speak() string
}

/*
class Dog
  method Speak() string //non-virtual
     return "Woof!"
*/
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

