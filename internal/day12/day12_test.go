package day12

import (
	"io"
	"strings"
	"testing"
)

var input = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

var input2 = `AAAA
BBCD
BBCC
EEEC`

var input3 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

var input4 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int(1930))
	e2 := any(int(1206))
	runTest(t, e1, e2, inputReader)
}

func Test_execute2(t *testing.T) {
	inputReader := strings.NewReader(input2)
	e1 := any(int(140))
	e2 := any(int(80))
	runTest(t, e1, e2, inputReader)
}

func Test_execute3(t *testing.T) {
	inputReader := strings.NewReader(input3)
	e1 := any(int(692))
	e2 := any(int(236))
	runTest(t, e1, e2, inputReader)
}

func Test_execute4(t *testing.T) {
	inputReader := strings.NewReader(input4)
	e1 := any(int(1184))
	e2 := any(int(368))
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
