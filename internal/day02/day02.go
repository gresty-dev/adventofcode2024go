package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

type dampener func(int) int

func Execute(input io.Reader) (aoc.Result, aoc.Result) {
	reports := readReports(input)
	result1 := aoc.NewResult(func() any {
		return countSafeReports(reports)
	})
	result2 := aoc.NewResult(func() any {
		return countSafeDampReports(reports)
	})

	return result1, result2
}

func countSafeReports(reports [][]int) int {
	count := 0
	for _, r := range reports {
		if isSafe(r) {
			count++
		}
	}
	return count
}

func countSafeDampReports(reports [][]int) int {
	count := 0
	for _, r := range reports {
		if isSafe(r) {
			count++
		} else {
			for s := range r {
				if isSafeWithDampening(r, func(i int) int { return skip(s, i) }) {
					count++
					break
				}
			}
		}
	}
	return count
}

func isSafe(report []int) bool {
	direction := signum(report[1] - report[0])
	for i := 1; i < len(report); i++ {
		if !isValid(direction, report[i], report[i-1]) {
			return false
		}
	}
	return true
}

func isSafeWithDampening(report []int, dampen dampener) bool {
	direction := signum(report[dampen(1)] - report[dampen(0)])
	for i := 1; i < len(report)-1; i++ {
		if !isValid(direction, report[dampen(i)], report[dampen(i-1)]) {
			return false
		}
	}
	return true
}

func skip(skip int, index int) int {
	if index >= skip {
		return index + 1
	}
	return index
}

func isValid(direction int, first int, second int) bool {
	diff := first - second
	if signum(diff) != direction {
		return false
	}
	if abs(diff) < 1 || abs(diff) > 3 {
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func signum(value int) int {
	if value == 0 {
		return 0
	} else if value > 0 {
		return 1
	} else {
		return -1
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func readReports(input io.Reader) [][]int {
	reports := [][]int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		report := []int{}
		for _, t := range tokens {
			n, _ := strconv.Atoi(t)
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	return reports
}
