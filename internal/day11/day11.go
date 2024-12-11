package day11

import (
	"bufio"
	"io"
	"math"
	"math/bits"
	"strconv"
	"strings"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type memoindex struct {
	stone  uint64
	blinks uint16
}

func Execute(input io.Reader) (Result, Result) {
	stones := readInput(input)

	memo := map[memoindex]uint64{}
	r1 := NewResult(func() any {
		return blinkAtStones(stones, 25, memo)
	})

	r2 := NewResult(func() any {
		return blinkAtStones(stones, 75, memo)
	})

	return r1, r2
}

func blinkAtStones(stones []uint64, blinks uint16, memo map[memoindex]uint64) uint64 {
	sum := uint64(0)
	for _, stone := range stones {
		sum += blink(stone, blinks, memo)
	}
	return sum
}

func blink(stone uint64, blinks uint16, memo map[memoindex]uint64) uint64 {
	if blinks == 0 {
		return 1
	}
	if cached, ok := memo[memoindex{stone, blinks}]; ok {
		return cached
	}

	var result uint64
	if stone == 0 {
		result = blink(1, blinks-1, memo)
	} else {
		digits := uint16(math.Log10(float64(stone)) + 1)
		if digits%2 == 0 {
			left, right := bits.Div64(0, stone, uint64(tenToThePowerOf(digits/2)))
			result = blink(left, blinks-1, memo) + blink(right, blinks-1, memo)
		} else {
			result = blink(stone*2024, blinks-1, memo)
		}
	}
	memo[memoindex{stone, blinks}] = result
	return result
}

func tenToThePowerOf(exp uint16) uint64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return 10
	}
	result := uint64(10)
	for i := uint16(2); i <= exp; i++ {
		result *= 10
	}
	return result
}

func readInput(input io.Reader) []uint64 {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()
	stones := []uint64{}
	for _, f := range strings.Fields(line) {
		stone, _ := strconv.ParseUint(f, 10, 64)
		stones = append(stones, stone)
	}
	return stones
}
