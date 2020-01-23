package main

import "fmt"

// class NamedObj
type NamedObj struct {
	Name string
}

//method show
func (n NamedObj) show() {
	Println(n.Name) // "n" is "this"
}

//class Rectangle
type Rectangle struct {
	NamedObj      //inheritance
	Width, Height float64
}

//override method show
func (r Rectangle) show() {
	Println("Rectangle ", r.Name) // "r" is "this"
}

func main() {
	var a = NamedObj{"Joe"}
	var b = Rectangle{NamedObj{"Richard"}, 10, 20}

	a.show("Hello")
	b.show("Hello")
}
