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
