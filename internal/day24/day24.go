package day24

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	. "go.gresty.dev/aoc2024/internal/lib"
)

var inputRegex = regexp.MustCompile(`(\w{3}): (\d)`)
var gateRegex = regexp.MustCompile(`(\w{3}) (OR|AND|XOR) (\w{3}) -> (\w{3})`)

type Gate struct {
	input  [2]string
	output int
	known  bool
	op     string
}

type System struct {
	bits  int
	gates map[string]*Gate
	// xInitial int64
	// yInitial int64
}

func NewSystem(bits int, gates *map[string]*Gate) System {
	s := System{
		bits:  bits,
		gates: *gates,
	}
	// s.xInitial = s.get("x")
	// s.yInitial = s.get("y")
	return s
}

func (s *System) set(name string, value int64) {
	mask := int64(1)
	for b := range s.bits {
		if gate, ok := s.gates[gateName(name, b)]; ok {
			if value&mask > 0 {
				gate.output = 1
			} else {
				gate.output = 0
			}
		}
	}
}

func (s *System) getGateValue(gateName string) int {
	gate, ok := s.gates[gateName]
	if !ok {
		return 0
	}
	if gate.known {
		return gate.output
	}
	i0 := s.getGateValue(gate.input[0])
	i1 := s.getGateValue(gate.input[1])
	switch gate.op {
	case "AND":
		gate.output = i0 & i1
	case "OR":
		gate.output = i0 | i1
	case "XOR":
		gate.output = i0 ^ i1
	}
	gate.known = true
	return gate.output
}

func (s *System) setGateValue(gateName string, value int) {
	if gate, ok := s.gates[gateName]; ok {
		gate.output = value
		gate.known = true
	}
}

func (s *System) get(name string) int64 {
	value := int64(0)
	for b := s.bits - 1; b >= 0; b-- {
		value = value << 1
		value += int64(s.getGateValue(gateName(name, b)))
	}
	return value
}

func (s *System) zero() {
	s.set("x", 0)
	s.set("y", 0)
	for n, g := range s.gates {
		if n[0] != 'x' && n[0] != 'y' {
			g.known = false
			g.output = 0
		}
	}
}

func (s System) findOutput(input0, input1, op string) (string, bool) {
	for n, g := range s.gates {
		if g.input[0] == input0 && g.input[1] == input1 && g.op == op {
			return n, true
		}
		if g.input[0] == input1 && g.input[1] == input0 && g.op == op {
			return n, true
		}
	}
	return fmt.Sprintf("No such gate %s %s %s", input0, op, input1), false
}

func (s *System) swap(first, second string) {
	s.gates[first], s.gates[second] = s.gates[second], s.gates[first]
}

func gateName(prefix string, bitNum int) string {
	return fmt.Sprintf("%s%02d", prefix, bitNum)
}

func Execute(input io.Reader) (Result, Result) {
	system := readInput(input)

	r1 := NewResult(func() any {
		return system.get("z")
	})
	r2 := NewResult(func() any {
		system.swap("hbs", "kfp")
		system.swap("z18", "dhq")
		system.swap("z22", "pdg")
		system.swap("z27", "jcp")
		verify(&system)
		verifyAdders(&system)
		return "dhq,hbs,jcp,kfp,pdg,z18,z22,z27"
	})
	return r1, r2
}

func ExecuteTest(input io.Reader) (Result, Result) {
	system := readInput(input)

	r1 := NewResult(func() any {
		return system.get("z")
	})
	r2 := NewResult(func() any {
		return 0
	})
	return r1, r2
}

func verify(system *System) {
	for b := range system.bits - 1 {
		for x := range 2 {
			for y := range 2 {
				last := (b == system.bits-1)
				system.zero()
				system.setGateValue(gateName("x", b), x)
				system.setGateValue(gateName("y", b), y)
				z := system.getGateValue(gateName("z", b))
				c := system.getGateValue(gateName("z", b+1))
				if z != x^y || (!last && c != x&y) || (last && c != 0) {
					fmt.Printf("Bad sum at bit %d - x=%d, y=%d, z=%d, c=%d\n", b, x, y, z, c)
				}
			}
		}
	}
}

func verifyAdders(system *System) {
	c := ""
	for bit := range system.bits - 1 {
		first := c == ""
		x := gateName("x", bit)
		y := gateName("y", bit)
		z := gateName("z", bit)
		h, ok := system.findOutput(x, y, "XOR")
		if !ok {
			panic(fmt.Sprintf("For bit %d, looking for x XOR y, %s", bit, h))
		}
		var out string
		if first {
			out = h
			if z != out {
				fmt.Printf("Bit %d mismatch - expected %s XOR %s -> %s but was %s\n", bit, x, y, z, out)
			}
		} else {
			out, ok = system.findOutput(h, c, "XOR")
			if !ok {
				panic(fmt.Sprintf("For bit %d, looking for c XOR h, %s", bit, out))
			}
			if z != out {
				fmt.Printf("Bit %d mismatch - expected h=%s XOR c=%s -> %s but was %s\n", bit, h, c, z, out)
			}
		}
		a, ok := system.findOutput(x, y, "AND")
		if !ok {
			panic(fmt.Sprintf("For bit %d, looking for x AND y, %s", bit, a))
		}
		if first {
			c = a
		} else {
			b, ok := system.findOutput(c, h, "AND")
			if !ok {
				panic(fmt.Sprintf("For bit %d, looking for c AND h, %s", bit, b))
			}
			c, ok = system.findOutput(a, b, "OR")
			if !ok {
				panic(fmt.Sprintf("For bit %d, looking for a AND b, %s", bit, c))
			}
		}
	}
}

func readInput(input io.Reader) System {
	scanner := bufio.NewScanner(input)
	gates := map[string]*Gate{}
	maxZ := -1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		name, gate := readInputLine(line)
		gates[name] = &gate
	}
	for scanner.Scan() {
		line := scanner.Text()
		name, gate := readGateLine(line)
		gates[name] = &gate
		if name[0] == 'z' {
			zIndex := Atoi(name[1:])
			if zIndex > maxZ {
				maxZ = zIndex
			}
		}
	}
	return NewSystem(maxZ+1, &gates)
}

func readInputLine(line string) (string, Gate) {
	matches := inputRegex.FindStringSubmatch(line)
	gate := Gate{
		output: Atoi(matches[2]),
		known:  true,
	}
	return matches[1], gate
}

func readGateLine(line string) (string, Gate) {
	matches := gateRegex.FindStringSubmatch(line)
	gate := Gate{
		input: [2]string{matches[1], matches[3]},
		op:    matches[2],
	}
	return matches[4], gate
}
