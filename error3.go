// Sample program to show how to implement a custom error type
// based on the json package in the standard library
package main

import (
   "fmt"
   "reflect"
)

// An UnmarshalTypeError describes a JSON value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
   Value string // description of json value
   Type reflect.Type // type of Go value it could not be assigned to
}

// Error implements the error interface.
func (e *UnmarshalTypeError) Error() string {
   return "json: cannot unmarshal " + e.value + "into Go value of type" + e.Type.String()
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// The argument to unmarshal must be non-nil pointer.
type InvalidUnmarshalError struct {
   Type reflect.Type
}

// Error implements the error interface
func (e *InvalidUnmarshalError) Error() string {
   if e.Type == nil {
      return "json: Unmarshal(nil)"
   }
   if e.Type.Kind() != reflect.Prt {
      return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
   }
   return "json: Unmarshal(nil " + e.Type.String() + ")"
}

// user is a type for use in the unmarshal call
type user struct {
   Name int
}

func main() {
   var u user
   err := Unmarshal([]byte(`{"name":"bill"}`), u) // Run with a value and pointer
   if err != nil {
      switch e := err.(type) {
      case *UnmarshalTypeError:
         fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
      case *InvalidUnmarshalError:
         fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.type)
      default:
         fmt.Println(err)
      }
      return
   }
   fmt.Println("Name:", u.Name)
}


