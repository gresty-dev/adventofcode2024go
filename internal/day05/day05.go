package day05

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

type ruletype map[int]struct{}
type updatetype []int

func Execute(input io.Reader) (aoc.Result, aoc.Result) {
	rules, updates := readInput(input)
	sumInOrder := 0
	sumReorder := 0
	for _, u := range updates {
		if inOrder(rules, u) {
			sumInOrder += middleValue(u)
		} else {
			reorder(u, rules)
			sumReorder += middleValue(u)
		}
	}
	return aoc.NewResult(func() any { return sumInOrder }), aoc.NewResult(func() any { return sumReorder })
}

func middleValue(update updatetype) int {
	return update[(len(update)-1)/2]
}

func inOrder(rules map[int]ruletype, update updatetype) bool {
	for i := 1; i < len(update); i++ {
		_, ok := rules[update[i-1]][update[i]]
		if !ok {
			return false
		}
	}
	return true
}

func reorder(update updatetype, rules map[int]ruletype) {
	sort.Slice(update, func(i, j int) bool {
		_, ok := rules[update[i]][update[j]]
		return ok // enough, cos every relation is defined one way or the other
	})
}

func readInput(input io.Reader) (map[int]ruletype, []updatetype) {
	rules := map[int]ruletype{}
	updates := []updatetype{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		a, _ := strconv.Atoi(line[:2])
		b, _ := strconv.Atoi(line[3:])
		rule := rules[a]
		if rule == nil {
			rule = ruletype{}
		}
		rule[b] = struct{}{}
		rules[a] = rule
	}

	for scanner.Scan() {
		line := scanner.Text()
		update := []int{}
		for _, token := range strings.Split(line, ",") {
			page, _ := strconv.Atoi(token)
			update = append(update, page)
		}
		updates = append(updates, update)
	}
	return rules, updates
}
