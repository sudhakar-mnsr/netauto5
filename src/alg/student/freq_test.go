package freq_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/freq"
)

// go test -run none -bench . -benchtime 3s

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	inp = buildText()
}

func buildText() []string {
	const n = 100
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, sentence)
	}
	for k, v := range out {
		out[k] = v * n
	}
	return s
}

var sentence = `The quick brown fox jumps over the lazy dog Stay hungry. Stay
foolish Keep going. Be all in Boldness be my friend Screw it, let’s do it My
life is my message Leave no stone unturned Dream big. Pray bigger`

var inp []string

var out = map[rune]int{
	'T': 1, 'q': 1, 'p': 2, '’': 1, 'i': 11, 'b': 4, 'w': 2, 'j': 1, 'B': 2,
	'L': 1, 'e': 20, 'v': 2, 'l': 7, ',': 1, 'h': 4, 'u': 5, 'f': 4, 's': 9,
	'g': 8, 'D': 1, 'P': 1, ' ': 37, 'z': 1, 'd': 5, '.': 3, 'c': 2, 'r': 9,
	'o': 11, 'm': 5, '\n': 2, 'x': 1, 'y': 8, 'S': 3, 'K': 1, 'k': 1, 'n': 10,
	't': 8, 'a': 8, 'M': 1,
}
