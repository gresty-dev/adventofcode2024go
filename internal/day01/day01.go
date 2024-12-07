package day01

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Execute(input io.Reader) (any, any) {
	list1, list2 := readIntSlicesFromInput(input)
	return part1(list1, list2), part2(list1, list2)
}

func part1(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += abs(list1[i] - list2[i])
	}
	return distance
}

func part2(list1 []int, list2 []int) int {
	counts := make(map[int]int)
	for _, v := range list2 {
		counts[v]++
	}
	similarity := 0
	for _, v := range list1 {
		similarity += v * counts[v]
	}
	return similarity
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readIntSlicesFromInput(input io.Reader) ([]int, []int) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		list1 = appendAsInt(list1, tokens[0])
		list2 = appendAsInt(list2, tokens[1])
	}
	return list1, list2
}

func appendAsInt(list []int, value string) []int {
	num, err := strconv.Atoi(value)
	check(err)
	list = append(list, num)
	return list
}
