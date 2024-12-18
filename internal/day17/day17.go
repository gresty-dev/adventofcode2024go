package day17

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type op func(*Computer, int)

var ops = []op{
	(*Computer).adv,
	(*Computer).bxl,
	(*Computer).bst,
	(*Computer).jnz,
	(*Computer).bxc,
	(*Computer).out,
	(*Computer).bdv,
	(*Computer).cdv,
}

type Computer struct {
	pc           int
	program      []int
	register     [3]int64
	output       []int
	registerSave [3]int64
}

func (c *Computer) reset() {
	c.register = c.registerSave
	c.pc = 0
	c.output = c.output[:0]
}

func (c *Computer) run() string {
	c.registerSave = c.register
	running := true
	for running {
		op := c.program[c.pc]
		arg := c.program[c.pc+1]
		c.pc += 2
		ops[op](c, arg)
		if c.pc >= len(c.program) {
			running = false
		}
	}
	return c.getOutput()
}

func (c Computer) getOutput() string {
	retval := ""
	for i, o := range c.output {
		if i > 0 {
			retval += ","
		}
		retval += strconv.Itoa(o)
	}
	return retval
}

func (c *Computer) adv(arg int) {
	c.register[0] = c.register[0] >> c.combo(arg)
}

func (c *Computer) bxl(arg int) {
	c.register[1] = c.register[1] ^ int64(arg)
}

func (c *Computer) bst(arg int) {
	c.register[1] = c.combo(arg) & 7
}

func (c *Computer) jnz(arg int) {
	if c.register[0] != 0 {
		c.pc = arg
	}
}

func (c *Computer) bxc(_ int) {
	c.register[1] = c.register[1] ^ c.register[2]
}

func (c *Computer) out(arg int) {
	c.output = append(c.output, int(c.combo(arg)&7))
}

func (c *Computer) bdv(arg int) {
	c.register[1] = c.register[0] >> c.combo(arg)
}

func (c *Computer) cdv(arg int) {
	c.register[2] = c.register[0] >> c.combo(arg)
}

func (c Computer) combo(arg int) int64 {
	if arg < 4 {
		return int64(arg)
	}
	return c.register[arg-4]
}

func Execute(input io.Reader) (Result, Result) {
	c := readInput(input)

	r1 := NewResult(func() any {
		return c.run()
	})
	r2 := NewResult(func() any {
		if answer, ok := findNextPartOfA(c, 0, len(c.program)-1); ok {
			return answer
		}
		return -1
	})
	return r1, r2
}

func findNextPartOfA(c Computer, a int64, pc int) (int64, bool) {
	if pc < 0 {
		return a, true
	}
	a = a << 3
	for i := 0; i < 8; i++ {
		c.reset()
		c.register[0] = a + int64(i)
		c.run()
		if c.output[0] == c.program[pc] {
			if answer, ok := findNextPartOfA(c, a+int64(i), pc-1); ok {
				return answer, true
			}
		}
	}
	return 0, false
}

var registerRegex = regexp.MustCompile(`Register .\: (\d+)`)

func readInput(input io.Reader) Computer {
	scanner := bufio.NewScanner(input)
	computer := Computer{}
	scanner.Scan()
	computer.register[0] = atoi64(registerRegex.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	computer.register[1] = atoi64(registerRegex.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	computer.register[2] = atoi64(registerRegex.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text()[9:], ",") {
		computer.program = append(computer.program, int(atoi64(v)))
	}
	return computer
}

func atoi64(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	panic(fmt.Sprintf("Cannot convert %s to integer", s))
}
