// go test -run none -bench . -benchtime 3s

// Tests to show how Data Oriented Design matters.
package caching

import "testing"

var fa int

// Capture the time it takes to perform a link list traversal.
func BenchmarkLinkListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	fa = a
}
