package day10

import (
	"image"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

func Execute(input io.Reader) (Result, Result) {
	topo := ReadGrid(input)
	topo.ForEachCell(func(loc image.Point, value byte) bool {
		topo.Set(loc, value-'0')
		return true
	})

	r1 := NewResult(func() any {
		return sumTrailheads(topo)
	})

	r2 := NewResult(func() any {
		return countUniqueTrails(topo)
	})

	return r1, r2
}

func sumTrailheads(topo Grid[byte]) int {
	endpoints := 0
	topo.ForEachCell(func(loc image.Point, value byte) bool {
		if value == 0 {
			endpoints += len(findNextStep(topo, loc, 0))
		}
		return true
	})
	return endpoints
}

func findNextStep(topo Grid[byte], cell image.Point, value byte) map[image.Point]Void {
	if value == 9 {
		return map[image.Point]Void{cell: Empty}
	}
	endpoints := map[image.Point]Void{}
	topo.ForEachNeighbour(cell, func(nb image.Point, nbval byte) bool {
		if nbval == value+1 {
			CombineSets(endpoints, findNextStep(topo, nb, nbval))
		}
		return true
	})
	return endpoints
}

func countUniqueTrails(topo Grid[byte]) int {
	heads := NewGrid[int](topo.Rows(), topo.Columns())
	cellsByHeight := [10][]image.Point{}
	topo.ForEachCell(func(loc image.Point, value byte) bool {
		topo.Set(loc, value)
		cellsByHeight[int(value)] = append(cellsByHeight[int(value)], loc)
		return true
	})

	for h := 9; h >= 0; h-- {
		for _, c := range cellsByHeight[h] {
			if h == 9 {
				heads.Set(c, 1)
			} else {
				heads.ForEachNeighbour(c, func(loc image.Point, value int) bool {
					if topo.Cell(loc) == byte(h+1) {
						heads.Set(c, heads.Cell(c)+value)
					}
					return true
				})
			}
		}
	}

	sum0 := 0
	for _, c := range cellsByHeight[0] {
		sum0 += heads.Cell(c)
	}

	return sum0
}
