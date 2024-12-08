package day08

import (
	"bufio"
	"image"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type pointSet map[image.Point]Void
type pointList []image.Point

type antinodeFinder func(image.Point, image.Point, image.Rectangle) pointSet

func Execute(input io.Reader) (Result, Result) {
	rect, antennae := readInput(input)
	result1 := NewResult(func() any {
		return len(findAntinodes(rect, antennae, antinodesForPair))
	})

	result2 := NewResult(func() any {
		return len(findAntinodes(rect, antennae, allAntinodesForPair))
	})

	return result1, result2
}

func findAntinodes(rect image.Rectangle, antennae map[byte]pointList, finder antinodeFinder) pointSet {
	antinodes := pointSet{}
	for _, v := range antennae {
		CombineSets(antinodes, antinodesForFrequency(rect, v, finder))
	}
	return antinodes
}

func antinodesForFrequency(rect image.Rectangle, nodes pointList, finder antinodeFinder) pointSet {
	antinodes := pointSet{}

	ForEachPair(nodes, func(a, b image.Point) {
		CombineSets(antinodes, finder(a, b, rect))
	})
	return antinodes
}

func antinodesForPair(a, b image.Point, rect image.Rectangle) pointSet {
	offset := b.Sub(a)
	antinodes := pointSet{}
	an1 := b.Add(offset)
	if an1.In(rect) {
		antinodes[an1] = Empty
	}
	an2 := a.Sub(offset)
	if an2.In(rect) {
		antinodes[an2] = Empty
	}
	return antinodes
}

func allAntinodesForPair(a, b image.Point, rect image.Rectangle) pointSet {
	offset := b.Sub(a)
	gcd := Gcd(offset.X, offset.Y)
	incr := offset.Div(gcd)

	antinodes := pointSet{}
	for an := a; an.In(rect); an = an.Add(incr) {
		antinodes[an] = Empty
	}
	for an := a; an.In(rect); an = an.Sub(incr) {
		antinodes[an] = Empty
	}

	return antinodes
}

func readInput(input io.Reader) (image.Rectangle, map[byte]pointList) {
	antennae := map[byte]pointList{}

	scanner := bufio.NewScanner(input)
	row := 0
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for col := range len(line) {
			a := line[col]
			if a == '.' {
				continue
			}
			p := image.Point{row, col}
			locs, ok := antennae[a]
			if !ok {
				locs = pointList{}
			}
			locs = append(locs, p)
			antennae[a] = locs
		}
		row++
	}
	gridrect := image.Rect(0, 0, row, cols)
	return gridrect, antennae
}
