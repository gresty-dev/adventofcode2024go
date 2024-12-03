package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var EXTRACT_MULS = regexp.MustCompile(`mul\(\d+,\d+\)`)
var EXTRACT_ENABLED_MULS = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don\'t\(\))`)
var EXTRACT_MUL_ARGS = regexp.MustCompile(`\d+`)

func main() {
	f, err := os.Open("input.txt")
	check(err)
	p1, p2 := execute(f)
	fmt.Println("Result: part1 =", p1, " part2 =", p2)
}

func execute(input io.Reader) (int, int) {
	program := readProgram(input)
	return addTheMuls(program), runProgram(program)
}

func execute1(input io.Reader) int {
	program := readProgram(input)
	return addTheMuls(program)
}

func execute2(input io.Reader) int {
	program := readProgram(input)
	return runProgram(program)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addTheMuls(program []string) int {
	sum := 0
	for _, line := range program {
		for _, mul := range extractMuls(line) {
			sum += executeMul(mul)
		}
	}
	return sum
}

func runProgram(program []string) int {
	sum := 0
	enabled := true
	var muls []string
	for _, line := range program {
		muls, enabled = extractEnabledMuls(line, enabled)
		for _, mul := range muls {
			sum += executeMul(mul)
		}
	}
	return sum
}

func extractMuls(line string) []string {
	return EXTRACT_MULS.FindAllString(line, -1)
}

func extractEnabledMuls(line string, enabled bool) ([]string, bool) {
	muls := []string{}
	for _, instr := range EXTRACT_ENABLED_MULS.FindAllString(line, -1) {
		switch {
		case instr == "do()":
			enabled = true
		case instr == "don't()":
			enabled = false
		default:
			if enabled {
				muls = append(muls, instr)
			}
		}
	}
	return muls, enabled
}

func executeMul(mul string) int {
	args := []int{}
	for _, a := range EXTRACT_MUL_ARGS.FindAllString(mul, 2) {
		intarg, _ := strconv.Atoi(a)
		args = append(args, intarg)
	}

	return args[0] * args[1]
}

func readProgram(input io.Reader) []string {
	program := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		program = append(program, line)
	}
	return program
}
