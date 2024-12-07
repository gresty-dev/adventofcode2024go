package day06

import (
	"strings"
	"testing"
)

var input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int(41))
	e2 := any(int(6))
	r1, r2 := Execute(inputReader)
	if r1.Answer() != e1 {
		t.Error("incorrect result for part 1: expected", e1, "got", r1.Answer())
	}
	if r2.Answer() != e2 {
		t.Error("incorrect result for part 2: expected", e2, "got", r2.Answer())
	}
}
