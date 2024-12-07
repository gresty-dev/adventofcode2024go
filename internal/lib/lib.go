package lib

import (
	"fmt"
	"io"
	"time"
)

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
	return fmt.Sprint(r.answer, " (", r.duration.Milliseconds(), " ms)")
}
