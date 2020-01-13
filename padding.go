// Sample program to show allignment, and how size varies
// and how to bring it down
package main

import (
   "fmt"
   "unsafe"
)

type flag struct {
        i64  int64
        f    bool
}
type example struct {
        flag    bool
        counter int16
        pi      float32
}
type pexample struct {
        pi      float32
        counter int16
        flag    bool
}
type big struct {
   flag1 bool
   counter1 int64
   flag2 bool
   counter2 int64
   flag3 bool
   counter3 int64
}
type pbig struct {
   counter1 int64
   counter2 int64
   counter3 int64
   flag1 bool
   flag2 bool
   flag3 bool
}
func main() {

	// Declare a variable of an anonymous type set
	// to its zero value.
        var f flag
	fmt.Printf("%d\n", unsafe.Sizeof(f))
	var e1 example
	// Display the value.
	fmt.Printf("%d\n", unsafe.Sizeof(e1))
        var e2 pexample
	// Display the values.
	fmt.Printf("%d\n", unsafe.Sizeof(e2))
        var b1 big
	fmt.Printf("%d\n", unsafe.Sizeof(b1))
        var b2 pbig
	fmt.Printf("%d\n", unsafe.Sizeof(b2))
        fmt.Printf("%p\n", &b2.counter1)
        fmt.Printf("%p\n", &b2.counter2)
        fmt.Printf("%p\n", &b2.counter3)
        fmt.Printf("%p\n", &b2.flag1)
        fmt.Printf("%p\n", &b2.flag2)
        fmt.Printf("%p\n", &b2.flag3)
}
