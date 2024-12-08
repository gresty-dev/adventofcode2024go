package lib

import (
	"fmt"
	"io"
	"time"
)

type Void struct{}

var Empty Void

type Solver func(io.Reader) (Result, Result)

type Result struct {
	duration time.Duration
	answer   any
}

func NewResult(operation func() any) Result {
	result := Result{}
	startTime := time.Now()
	result.answer = operation()
	result.duration = time.Since(startTime)
	return result
}

func (r Result) Answer() any {
	return r.answer
}

func (r Result) String() string {
	return fmt.Sprint(r.answer, " (", r.duration.Microseconds(), " us)")
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func ForEachPair[T any](set []T, operation func(T, T)) {
	for i := 0; i < len(set)-1; i++ {
		for j := i + 1; j < len(set); j++ {
			operation(set[i], set[j])
		}
	}
}

func CombineSets[T comparable](m1, m2 map[T]Void) {
	for a := range m2 {
		m1[a] = Empty
	}
}
