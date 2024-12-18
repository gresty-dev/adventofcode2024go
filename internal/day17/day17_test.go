package day17

import (
	"io"
	"strings"
	"testing"
)

var input = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var input2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func Test_execute(t *testing.T) {
	inputReader := strings.NewReader(input)
	e1 := any("4,6,3,5,6,3,5,2,1,0")
	e2 := any(int64(29328))
	runTest(t, e1, e2, inputReader)
}

func Test_execute2(t *testing.T) {
	inputReader := strings.NewReader(input2)
	e1 := any("5,7,3,0")
	e2 := any(int64(117440))
	runTest(t, e1, e2, inputReader)
}

func Test_ops1(t *testing.T) {
	c := Computer{
		program:  []int{2, 6},
		register: [3]int64{0, 0, 9},
	}
	c.run()
	assertRegister(t, "B", 1, c.register[1])
}

func Test_ops2(t *testing.T) {
	c := Computer{
		program:  []int{5, 0, 5, 1, 5, 4},
		register: [3]int64{10, 0, 0},
	}
	assertOutput(t, "0,1,2", c.run())
}

func Test_ops3(t *testing.T) {
	c := Computer{
		program:  []int{0, 1, 5, 4, 3, 0},
		register: [3]int64{2024, 0, 0},
	}
	assertOutput(t, "4,2,5,6,7,7,7,7,3,1,0", c.run())
	assertRegister(t, "A", 0, c.register[0])
}

func Test_ops4(t *testing.T) {
	c := Computer{
		program:  []int{1, 7},
		register: [3]int64{0, 29, 0},
	}
	c.run()
	assertRegister(t, "B", 26, c.register[1])
}

func Test_ops5(t *testing.T) {
	c := Computer{
		program:  []int{4, 0},
		register: [3]int64{0, 2024, 43690},
	}
	c.run()
	assertRegister(t, "B", 44354, c.register[1])
}

func assertOutput(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("Expected output to be %s but was %s", expected, actual)
	}
}

func assertRegister(t *testing.T, register string, expected, actual int64) {
	if actual != expected {
		t.Errorf("Expected register %s to be %d but was %d", register, expected, actual)
	}
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
