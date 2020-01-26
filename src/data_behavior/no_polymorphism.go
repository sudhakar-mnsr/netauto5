// Sample program that could benefit from polymorphic behavior with 
// interfaces
package main

import "fmt"

// File defines a system file
type file struct {
   name string
}

// read implements the reader interface for a file
func (file) read(b []byte) (int, error) {
   s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
   copy(b,s)
   return len(s), nil
}

type pipe struct {
   name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
   s := `{name: "bill", title: "developer"}`
   copy(b, s)
   return len(s), nil
}

func main() {
   f := file{"data.Json"}
   retrieveFile(f)
}

func retrieveFile(f file) error {
   data := make([]byte, 100)
   
   len, err := f.read(data)
   if err != nil {
      return err
   }

   fmt.Println(string(data[:len]))
   return nil
}

func retrievePipe(p pipe) error {
   data := make([]byte, 100)

   len, err := p.read(data)
   if err != nil {
      return err
   }

   fmt.Println(string(data[:len]))
   return nil
}
