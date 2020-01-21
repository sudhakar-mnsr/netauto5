package freq_test

import (
	"testing"

	freq "student/freq"
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
