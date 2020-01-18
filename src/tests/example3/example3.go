// Sample program that implements a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/go/testing/tests/example4/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
