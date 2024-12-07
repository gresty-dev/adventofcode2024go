package day07

import (
	"strings"
	"testing"
)

var input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int64(3749))
	e2 := any(int64(11387))
	r1, r2 := Execute(inputReader)
	if r1.Answer() != e1 {
		t.Error("incorrect result for part 1: expected", e1, "got", r1.Answer())
	}
	if r2.Answer() != e2 {
		t.Error("incorrect result for part 2: expected", e2, "got", r2.Answer())
	}
}
