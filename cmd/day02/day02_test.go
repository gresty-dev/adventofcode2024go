package main

import (
	"strings"
	"testing"
)

var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	part1, part2 := execute(inputReader)
	if part1 != 2 {
		t.Error("incorrect result for part 1: expected 2, got ", part1)
	}
	if part2 != 4 {
		t.Error("incorrect result for part 2: expected 4, got ", part2)
	}
}
