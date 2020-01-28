// Sample program demonstrating interface composition.
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// =============================================================================

// Data is the structure of the data we are copying.
type Data struct {
	Line string
}
// =============================================================================

// Puller declares behavior for pulling data.
type Puller interface {
   Pull(d *Data) error
}
type Storer interface {
   Store(d *Data) error
}

// PullStorer declares behavior for both pulling and storing.
type PullStorer interface {
   Puller
   Storer
}
// =============================================================================
// Xenia is a system we need to pull data from.
type Xenia struct {
   Host string
   Timeout time.Duration
}

// Pull knows how to pull data out of Xenia.
func (*Xenia) Pull(d *Data) error {
   switch rand.Intn(10) {
      case 1,9:
         return io.EOF
      case 5:
         return errors.New("Error reading data from Xenia")
      default:
         d.Line = "Data"
         fmt.Println("In:", d.Line)
         return nil
   }
}  

// Pillar is a system we need to store data into.
type Pillar struct {
   Host string
   Timeout time.Duration
}

// Store knows how to store data into Pillar
func (*Pillar) Store(d *Data) error {
   fmt.Println("Out:", d.Line)
   return nil
}

// =============================================================================
// System wraps Xenia and Pillar together into a single system.
type System struct {
   Xenia
   Pillar
}
// =============================================================================

