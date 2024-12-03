package main

import (
	"strings"
	"testing"
)

var input1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
var input2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func Test_execute1(t *testing.T) {
	inputReader := strings.NewReader(input1)
	result := execute1(inputReader)
	if result != 161 {
		t.Error("incorrect result for part 1: expected 161, got ", result)
	}
}

func Test_execute2(t *testing.T) {
	inputReader := strings.NewReader(input2)
	result := execute2(inputReader)
	if result != 48 {
		t.Error("incorrect result for part 2: expected 48, got ", result)
	}
}
