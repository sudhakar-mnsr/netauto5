// Sample program to show how maps are reference types.
package main

import "fmt"

func main() {

	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a function to perform some mutation.
	double(scores, "anna")

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}
