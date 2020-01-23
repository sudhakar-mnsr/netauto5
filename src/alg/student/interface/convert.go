package main

import (
   "fmt"
   "strconv"
)

type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
//    case float32:
//        return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}

func main() {
   fmt.Println(ToString(1))
}
