package day09

import (
	"io"
	"strings"
	"testing"
)

var input = `2333133121414131402`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int64(1928))
	e2 := any(int64(2858))
	runTest(t, e1, e2, inputReader)
}

func runTest(t *testing.T, e1, e2 any, reader io.Reader) {
	r1, r2 := Execute(reader)
	if r1.Answer() != e1 {
		t.Error("incorrect result for part 1: expected", e1, "got", r1.Answer())
	}
	if r2.Answer() != e2 {
		t.Error("incorrect result for part 2: expected", e2, "got", r2.Answer())
	}
}
