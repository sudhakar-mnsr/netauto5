// Sample program to show how to declare and initialize anonymous
// struct types.
package main

import "fmt"

type bill struct {
        flag    bool
        counter int16
        pi      float32
}
type alice struct {
        flag    bool
        counter int16
        pi      float32
}

func main() {
        var a alice 
        var b bill 
        fmt.Println(a,b)
        a=alice(b)
	// Declare a variable of an anonymous type set
	// to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
        var c alice
        c = e2
	fmt.Printf("c address: %p\n", &c)
	fmt.Printf("e2 address: %p\n", &e2)
}
