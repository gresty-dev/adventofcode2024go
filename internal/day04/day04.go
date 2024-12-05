package day04

import (
	"io"
	"strings"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

func Execute(input io.Reader) (int, int) {
	grid := aoc.ReadGrid(input)
	return countXmas(grid), countMasInX(grid)
}

func countMasInX(grid aoc.Grid) int {
	count := 0

	grid.ForEachCell(func(r, c int, v byte) bool {
		if !grid.NearBoundary(r, c, 0) && v == 'A' && masInXFound(grid, r, c) {
			count++
		}
		return true
	})

	return count
}

func masInXFound(grid aoc.Grid, r int, c int) bool {
	var sb strings.Builder
	sb.WriteByte(grid.Cell(r-1, c-1))
	sb.WriteByte(grid.Cell(r+1, c+1))
	sb.WriteByte(grid.Cell(r-1, c+1))
	sb.WriteByte(grid.Cell(r+1, c-1))
	w := sb.String()
	return w == "MSMS" || w == "SMMS" || w == "MSSM" || w == "SMSM"
}

func countXmas(grid aoc.Grid) int {
	count := 0
	grid.ForEachCell(func(r, c int, v byte) bool {
		if v != 'X' {
			return true
		}
		for _, d := range aoc.Directions {
			if grid.WordAt(r, c, 4, d) == "XMAS" {
				count++
			}
		}
		return true
	})
	return count
}
