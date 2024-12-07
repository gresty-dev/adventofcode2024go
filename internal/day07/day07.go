package day07

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type equation struct {
	solution int64
	factors  []int64
}

func Execute(input io.Reader) (any, any) {
	equations := readInput(input)

	var sum1 int64
	for _, equation := range equations {
		if calculate(equation.solution, 0, equation.factors) {
			sum1 += equation.solution
		}
	}

	var sum2 int64
	for _, equation := range equations {
		if calculateWithConcat(equation.solution, 0, equation.factors) {
			sum2 += equation.solution
		}
	}

	return sum1, sum2
}

func calculate(target int64, sofar int64, factors []int64) bool {
	if sofar == 0 {
		return calculate(target, factors[0], factors[1:])
	}
	if sofar > target {
		return false
	}
	if len(factors) == 0 {
		return sofar == target
	}
	return calculate(target, sofar*factors[0], factors[1:]) ||
		calculate(target, sofar+factors[0], factors[1:])
}

func calculateWithConcat(target int64, sofar int64, factors []int64) bool {
	if sofar == 0 {
		return calculateWithConcat(target, factors[0], factors[1:])
	}
	if sofar > target {
		return false
	}
	if len(factors) == 0 {
		return sofar == target
	}
	return calculateWithConcat(target, sofar*factors[0], factors[1:]) ||
		calculateWithConcat(target, sofar+factors[0], factors[1:]) ||
		calculateWithConcat(target, concat(sofar, factors[0]), factors[1:])
}

func concat(a, b int64) int64 {
	result, _ := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
	return result
}

func readInput(input io.Reader) []equation {
	equations := []equation{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		solution, _ := strconv.ParseInt(tokens[0][:len(tokens[0])-1], 10, 64)
		factors := []int64{}
		for i := 1; i < len(tokens); i++ {
			factor, _ := strconv.ParseInt(tokens[i], 10, 64)
			factors = append(factors, factor)
		}
		equations = append(equations, equation{solution, factors})
	}
	return equations
}
