package day04

import (
	"image"
	"io"
	"strings"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

func Execute(input io.Reader) (aoc.Result, aoc.Result) {
	grid := aoc.ReadGrid(input)
	result1 := aoc.NewResult(func() any {
		return countXmas(grid)
	})
	result2 := aoc.NewResult(func() any {
		return countMasInX(grid)
	})
	return result1, result2
}

func countMasInX(grid aoc.Grid[byte]) int {
	count := 0

	grid.ForEachCell(func(loc image.Point, v byte) {
		if !grid.NearBoundary(loc, 0) && v == 'A' && masInXFound(grid, loc) {
			count++
		}
	})

	return count
}

func masInXFound(grid aoc.Grid[byte], loc image.Point) bool {
	var sb strings.Builder
	sb.WriteByte(grid.Cell(loc.Add(aoc.UpLeft.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.DownRight.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.UpRight.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.DownLeft.Point)))
	w := sb.String()
	return w == "MSMS" || w == "SMMS" || w == "MSSM" || w == "SMSM"
}

func countXmas(grid aoc.Grid[byte]) int {
	count := 0
	grid.ForEachCell(func(loc image.Point, v byte) {
		if v == 'X' {
			for _, d := range aoc.Directions {
				if string(grid.WordAt(loc, 4, d)) == "XMAS" {
					count++
				}
			}
		}
	})
	return count
}
