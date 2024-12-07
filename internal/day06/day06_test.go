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
	r1, r2 := Execute(inputReader)
	if r1 != 41 {
		t.Error("incorrect result for part 1: expected 41, got ", r1)
	}
	if r2 != 6 {
		t.Error("incorrect result for part 2: expected 6, got ", r2)
	}
}
