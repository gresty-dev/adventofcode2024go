package day10

import (
	"io"
	"strings"
	"testing"
)

var input1 = `0123
1234
8765
9876`

var input2 = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Test_execute1(t *testing.T) {
	inputReader := strings.NewReader(input1)
	e1 := any(int(1))
	e2 := any(int(16))
	runTest(t, e1, e2, inputReader)
}

func Test_execute2(t *testing.T) {
	inputReader := strings.NewReader(input2)
	e1 := any(int(36))
	e2 := any(int(81))
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
