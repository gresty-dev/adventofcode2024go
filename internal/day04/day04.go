package day04

import (
	"image"
	"io"
	"strings"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

func Execute(input io.Reader) (any, any) {
	grid := aoc.ReadGrid(input)
	return countXmas(grid), countMasInX(grid)
}

func countMasInX(grid aoc.Grid) int {
	count := 0

	grid.ForEachCell(func(loc image.Point, v byte) bool {
		if !grid.NearBoundary(loc, 0) && v == 'A' && masInXFound(grid, loc) {
			count++
		}
		return true
	})

	return count
}

func masInXFound(grid aoc.Grid, loc image.Point) bool {
	var sb strings.Builder
	sb.WriteByte(grid.Cell(loc.Add(aoc.UpLeft.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.DownRight.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.UpRight.Point)))
	sb.WriteByte(grid.Cell(loc.Add(aoc.DownLeft.Point)))
	w := sb.String()
	return w == "MSMS" || w == "SMMS" || w == "MSSM" || w == "SMSM"
}

func countXmas(grid aoc.Grid) int {
	count := 0
	grid.ForEachCell(func(loc image.Point, v byte) bool {
		if v != 'X' {
			return true
		}
		for _, d := range aoc.Directions {
			if grid.WordAt(loc, 4, d) == "XMAS" {
				count++
			}
		}
		return true
	})
	return count
}
