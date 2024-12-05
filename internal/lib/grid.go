package lib

import (
	"bufio"
	"image"
	"io"
)

type cellhandler func(row, col int, value byte) bool

// X maps to row, Y maps to column!!
var Directions = []image.Point{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1},
}

type Grid struct {
	cells []string
}

func ReadGrid(input io.Reader) Grid {
	grid := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	return Grid{grid}
}

func (g Grid) Cell(row, col int) byte {
	return g.cells[row][col]
}

func (g Grid) Columns() int {
	return len(g.cells[0])
}

func (g Grid) Rows() int {
	return len(g.cells)
}

func (g Grid) ForEachCell(percell cellhandler) {
	for r := range g.Rows() {
		for c := range g.Columns() {
			if !percell(r, c, g.Cell(r, c)) {
				break
			}
		}
	}
}

func (g Grid) NearBoundary(row, col, dist int) bool {
	return row <= dist || col <= dist ||
		row >= g.Rows()-dist-1 || col >= g.Columns()-dist-1
}

func (g Grid) InGrid(row, col int) bool {
	return row >= 0 && row < g.Rows() && col >= 0 && col < g.Columns()
}

func (g Grid) WordAt(row, col, length int, dir image.Point) string {
	limit := dir.Mul(length - 1)
	if !g.InGrid(row+limit.X, col+limit.Y) {
		return ""
	}

	word := []byte{}
	for range length {
		word = append(word, g.Cell(row, col))
		row += dir.X
		col += dir.Y
	}
	return string(word)
}
