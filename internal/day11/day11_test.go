package day11

import (
	"io"
	"strings"
	"testing"
)

var input = `125 17`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(uint64(55312))
	e2 := any(uint64(0))
	runTest(t, e1, e2, inputReader)
}

func runTest(t *testing.T, e1, _ any, reader io.Reader) {
	r1, _ := Execute(reader)
	if r1.Answer() != e1 {
		t.Error("incorrect result for part 1: expected", e1, "got", r1.Answer())
	}
}
