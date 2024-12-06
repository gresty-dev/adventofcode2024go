package day05

import (
	"strings"
	"testing"
)

var input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	r1, r2 := Execute(inputReader)
	if r1 != 143 {
		t.Error("incorrect result for part 1: expected 143, got ", r1)
	}
	if r2 != 123 {
		t.Error("incorrect result for part 2: expected 123, got ", r2)
	}
}
