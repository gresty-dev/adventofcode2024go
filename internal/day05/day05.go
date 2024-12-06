package day05

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type ruletype map[int]struct{}
type updatetype []int

func Execute(input io.Reader) (int, int) {
	rules, updates := readInput(input)
	sumInOrder := 0
	sumReorder := 0
	for _, u := range updates {
		if validateUpdate(rules, u) {
			sumInOrder += middleValue(u)
		} else {
			reorder(u, rules)
			sumReorder += middleValue(u)
		}
	}
	return sumInOrder, sumReorder
}

func middleValue(update updatetype) int {
	return update[(len(update)-1)/2]
}

func validateUpdate(rules map[int]ruletype, update updatetype) bool {
	for i := 1; i < len(update); i++ {
		if isInOrder(update[i-1], update[i], rules) {
			continue
		}
		return false
	}
	return true
}

func isInOrder(first, second int, rules map[int]ruletype) bool {
	_, ok := rules[second][first]
	return !ok
}

func reorder(update updatetype, rules map[int]ruletype) {
	sort.Slice(update, func(i, j int) bool {
		first := update[i]
		second := update[j]
		_, ok := rules[first][second]
		if ok {
			return true
		}
		_, ok = rules[second][first]
		return !ok
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
