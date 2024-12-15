package day13

import (
	"io"
	"strings"
	"testing"
)

var input = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int64(480))
	e2 := any(int64(875318608908))
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
