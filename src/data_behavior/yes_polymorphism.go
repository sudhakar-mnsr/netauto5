// Sample program to show how polymorphic behavior with interfaces.

package main

import "fmt"

// reader is an interface that defines the act of reading data
type reader interface {
   read(b []byte) (int, error)
}

type file struct {
   name string
}

func (file) read(b []byte) (int, error) {
   s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
   copy(b, s)
   return len(s), nil
}

func main() {
// create two values one of type file and one of type pipe
f := file{"data.json"}

// Call retrieve function for each concrete type
retrieve(f)
}

// retrieve can read any device and process the data
func retrieve(r reader) error {
   data := make([]byte, 100)

   len, err := r.read(data)
   if err != nil {
      return err
   }

   fmt.Println(string(data[:len]))
   return nil
}
