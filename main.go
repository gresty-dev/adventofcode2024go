package main

import (
	"fmt"
	"os"
	"strconv"

	day01 "go.gresty.dev/aoc2024/internal/day01"
	day02 "go.gresty.dev/aoc2024/internal/day02"
	day03 "go.gresty.dev/aoc2024/internal/day03"
	day04 "go.gresty.dev/aoc2024/internal/day04"
	day05 "go.gresty.dev/aoc2024/internal/day05"
	day06 "go.gresty.dev/aoc2024/internal/day06"
	day07 "go.gresty.dev/aoc2024/internal/day07"
	lib "go.gresty.dev/aoc2024/internal/lib"
)

var solvers = map[int]lib.Solver{
	1: day01.Execute,
	2: day02.Execute,
	3: day03.Execute,
	4: day04.Execute,
	5: day05.Execute,
	6: day06.Execute,
	7: day07.Execute,
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

	daynum, err := strconv.Atoi(args[0])
	if err != nil {
		usage()
	}

	inputFile := fmt.Sprintf("internal/day%02d/input.txt", daynum)
	solve(solvers[daynum], inputFile)
}

func usage() {
	fmt.Println("Usage: ", os.Args[0], "DAYNUM")
}

func solve(solution lib.Solver, filename string) {
	f, err := os.Open(filename)
	check(err)
	p1, p2 := solution(f)
	fmt.Println("Part1: ", p1.String())
	fmt.Println("Part2: ", p2.String())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
