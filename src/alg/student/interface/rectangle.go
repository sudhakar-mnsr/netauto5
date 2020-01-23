package main

import . "fmt"

type Rectangle struct {
	Name          string
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var a Rectangle
	var b = Rectangle{"I'm b.", 10, 20}
	var c = Rectangle{Height: 12, Width: 14}
        var d = b.Area()

	Println(a)
	Println(b)
	Println(c)
	Println(d)
}
