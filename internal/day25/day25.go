package day25

import (
	"bufio"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type Node struct {
	children []*Node
}

func (n *Node) add(key []int) {
	if n.children[key[0]] == nil {
		n.children[key[0]] = &(Node{make([]*Node, 6)})
	}
	if len(key) > 1 {
		n.children[key[0]].add(key[1:])
	}
}

func (n Node) countMatches(key []int) int {
	if len(key) == 0 {
		return 1
	}
	sum := 0
	for i := 0; i <= 5-key[0]; i++ {
		if n.children[i] != nil {
			sum += n.children[i].countMatches(key[1:])
		}
	}
	return sum
}

type KeyLock struct {
	pins []int
}

func Execute(input io.Reader) (Result, Result) {
	keys, locks := readInput(input)

	r1 := NewResult(func() any {
		return fitKeysToLocks(keys, locks)
	})
	r2 := NewResult(func() any {
		return 0
	})
	return r1, r2
}

func fitKeysToLocks(keys []KeyLock, locks *Node) int {
	sum := 0
	for _, k := range keys {
		sum += locks.countMatches(k.pins)
	}
	return sum
}

func readInput(input io.Reader) ([]KeyLock, *Node) {
	scanner := bufio.NewScanner(input)
	running := true
	keys := []KeyLock{}
	locks := Node{make([]*Node, 6)}
	for running {
		keylock := [5]int{-1, -1, -1, -1, -1}
		isKey := false
		for i := range 7 {
			scanner.Scan()
			line := scanner.Text()
			if i == 0 {
				isKey = line[0] == '.'
			}
			for j := range 5 {
				if line[j] == '#' {
					keylock[j]++
				}
			}
		}
		if isKey {
			keys = append(keys, KeyLock{keylock[0:]})
		} else {
			locks.add(keylock[0:])
		}
		running = scanner.Scan()
	}
	return keys, &locks
}
