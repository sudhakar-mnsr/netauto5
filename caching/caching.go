// Package caching provides code to show why Data Oriented Design matters. How
// data layouts matter more to performance than algorithm efficiency.
package caching

import "fmt"

// Create a square matrix of 2 meg by 2 meg.
const (
	rows = 2 * 1024
	cols = 2 * 1024
)

// matrix represents a matrix with a large number of
// columns per row.
var matrix [rows][cols]byte

// data represents a data node for our linked list.
type data struct {
	v byte
	p *data
}

// list points to the head of the list.
var list *data

