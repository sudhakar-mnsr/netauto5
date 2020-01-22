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
