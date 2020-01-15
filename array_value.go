package main

import "fmt"

func main() {
    friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
    fmt.Printf("Bfr[%s] : ", friends[1])

    for i, v := range friends {
            friends[1] = "Jack"

            if i == 1 {
                    fmt.Printf("v[%p -> %s]\n", &v, v)
            }
    }
    fmt.Println("Outside end", friends)
    fmt.Printf("Address of friends[1]: %p\n", &friends[1]) 
}
