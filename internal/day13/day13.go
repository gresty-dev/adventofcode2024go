package day13

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	. "go.gresty.dev/aoc2024/internal/lib"
)

var buttonRegex = regexp.MustCompile(`Button [A|B]: X\+(\d+), Y\+(\d+)`)
var prizeRegex = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

type Machine struct {
	ax, ay, bx, by int64
	px, py         int64
}

// axA + bxB = px
// ayA + byB = py
func (m Machine) solve() (int64, int64, bool) {
	anum := m.px*m.by - m.bx*m.py
	aden := m.ax*m.by - m.bx*m.ay
	if aden != 0 && anum%aden == 0 {
		a := anum / aden
		bnum := m.px - m.ax*a
		bden := m.bx
		if bden != 0 && bnum%bden == 0 {
			b := bnum / bden
			return a, b, true
		}
	}
	return 0, 0, false
}

func Execute(input io.Reader) (Result, Result) {
	machines := readInput(input)

	r1 := NewResult(func() any {
		return winAllPrizes(machines, 0)
	})
	r2 := NewResult(func() any {
		return winAllPrizes(machines, 10_000_000_000_000)
	})
	return r1, r2
}

func winAllPrizes(machines []Machine, offset int64) int64 {
	sum := int64(0)
	for _, m := range machines {
		offsetMachine := m
		offsetMachine.px += offset
		offsetMachine.py += offset
		if a, b, ok := offsetMachine.solve(); ok {
			sum += 3*a + b
		}
	}
	return sum
}

func readInput(input io.Reader) []Machine {
	scanner := bufio.NewScanner(input)
	machines := []Machine{}
	var latest Machine
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 9 && line[:9] == "Button A:" {
			latest = Machine{}
			matches := buttonRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			latest.ax = int64(x)
			latest.ay = int64(y)
		} else if len(line) >= 9 && line[:9] == "Button B:" {
			matches := buttonRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			latest.bx = int64(x)
			latest.by = int64(y)
		} else if len(line) >= 6 && line[:6] == "Prize:" {
			matches := prizeRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			latest.px = int64(x)
			latest.py = int64(y)
			machines = append(machines, latest)
		}
	}
	return machines
}
