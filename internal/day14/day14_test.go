package day14

import (
	"io"
	"strings"
	"testing"
)

var input = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any(int64(12))
	e2 := any(int64(0))
	runTest(t, e1, e2, inputReader)
}

func runTest(t *testing.T, e1, e2 any, reader io.Reader) {
	r1, r2 := ExecuteTest(reader)
	if r1.Answer() != e1 {
		t.Error("incorrect result for part 1: expected", e1, "got", r1.Answer())
	}
	if r2.Answer() != e2 {
		t.Error("incorrect result for part 2: expected", e2, "got", r2.Answer())
	}
}
