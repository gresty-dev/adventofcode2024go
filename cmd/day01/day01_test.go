package main

import (
	"strings"
	"testing"
)

var input = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	part1, part2 := execute(inputReader)
	if part1 != 11 {
		t.Error("incorrect result for part 1: expected 11, got ", part1)
	}
	if part2 != 31 {
		t.Error("incorrect result for part 2: expected 31, got ", part2)
	}
}
