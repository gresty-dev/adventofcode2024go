package day04

import (
	"fmt"
	"strings"
	"testing"
)

var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	r1, r2 := Execute(inputReader)
	if r1 != 18 {
		t.Error("incorrect result for part 1: expected 18, got ", r1)
	}
	if r2 != 9 {
		t.Error("incorrect result for part 2: expected 9, got ", r2)
	}
}

func Test_small(t *testing.T) {
	input := `MDM
DAD
SDS`

	fmt.Println(input)
	_, r2 := Execute(strings.NewReader(input))
	if r2 != 1 {
		t.Error("incorrect result for part 2: expected 1, got ", r2)
	}

}
