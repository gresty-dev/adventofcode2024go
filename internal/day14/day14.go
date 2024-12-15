package day14

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"regexp"
	"strconv"

	. "go.gresty.dev/aoc2024/internal/lib"
)

var robotRegex = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

type Robot struct {
	x, y   int
	dx, dy int
}

func Execute(input io.Reader) (Result, Result) {
	robots := readInput(input)

	r1 := NewResult(func() any {
		return runSim(robots, 100, 101, 103)
	})
	r2 := NewResult(func() any {
		return findChristmasTree(robots, 101, 103)
	})
	return r1, r2
}

func ExecuteTest(input io.Reader) (Result, Result) {
	robots := readInput(input)

	r1 := NewResult(func() any {
		return runSim(robots, 100, 11, 7)
	})
	r2 := NewResult(func() any {
		return int64(0)
	})
	return r1, r2
}

func runSim(robots []Robot, time int, xlen, ylen int) int64 {
	xmid := xlen / 2
	ymid := ylen / 2
	quadCounts := [4]int64{}
	for _, r := range robots {
		x := pmod(r.x+time*r.dx, xlen)
		y := pmod(r.y+time*r.dy, ylen)
		if x < xmid {
			if y < ymid {
				quadCounts[0]++
			} else if y > ymid {
				quadCounts[2]++
			}
		} else if x > xmid {
			if y < ymid {
				quadCounts[1]++
			} else if y > ymid {
				quadCounts[3]++
			}
		}
	}
	return quadCounts[0] * quadCounts[1] * quadCounts[2] * quadCounts[3]
}

func findChristmasTree(robots []Robot, xlen, ylen int) int64 {
	xmid := xlen / 2
	ymid := ylen / 2
	safetyFactor := int64(10_000_000_000_000)
	time := 0
	for time = 1; safetyFactor >= 99235800; time++ { // found this by trial and error!
		quadCounts := [4]int64{}
		for i, r := range robots {
			x := pmod(r.x+r.dx, xlen)
			y := pmod(r.y+r.dy, ylen)
			if x < xmid {
				if y < ymid {
					quadCounts[0]++
				} else if y > ymid {
					quadCounts[2]++
				}
			} else if x > xmid {
				if y < ymid {
					quadCounts[1]++
				} else if y > ymid {
					quadCounts[3]++
				}
			}
			r.x = x
			r.y = y
			robots[i] = r
		}
		safetyFactor = quadCounts[0] * quadCounts[1] * quadCounts[2] * quadCounts[3]
	}

	dumpRobots(robots, time, xlen, ylen)
	return safetyFactor
}

func dumpRobots(robots []Robot, time int, xlen, ylen int) {
	grid := NewGrid[string](ylen, xlen)
	grid.ForEachCell(func(loc image.Point, _ string) {
		grid.Set(loc, " ")
	})

	for _, r := range robots {
		grid.Set(image.Point{r.y, r.x}, "X")
	}

	fmt.Println("After ", time-1, " seconds")
	grid.Dump()
}

func readInput(input io.Reader) []Robot {
	scanner := bufio.NewScanner(input)
	robots := []Robot{}
	for scanner.Scan() {
		matches := robotRegex.FindStringSubmatch(scanner.Text())
		robot := Robot{
			x:  atoi(matches[1]),
			y:  atoi(matches[2]),
			dx: atoi(matches[3]),
			dy: atoi(matches[4]),
		}
		robots = append(robots, robot)
	}
	return robots
}

func atoi(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	panic(fmt.Sprintf("Cannot convert %s to integer", s))
}

// Positive modulo, returns non negative solution to x % d
func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}
