package lib

import (
	"bufio"
	"image"
	"io"
)

type cellhandler[T comparable] func(loc image.Point, value T) bool

type Grid[T comparable] struct {
	cells [][]T
}

func NewGrid[T comparable](rows, columns int) Grid[T] {
	grid := make([][]T, rows)
	for r := range rows {
		grid[r] = make([]T, columns)
	}
	return Grid[T]{grid}
}

func ReadGrid(input io.Reader) Grid[byte] {
	grid := [][]byte{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	return Grid[byte]{grid}
}

func (g Grid[T]) Cell(loc image.Point) T {
	return g.cells[loc.X][loc.Y]
}

func (g *Grid[T]) Set(loc image.Point, value T) {
	g.cells[loc.X][loc.Y] = value
}

func (g Grid[T]) Columns() int {
	return len(g.cells[0])
}

func (g Grid[T]) Rows() int {
	return len(g.cells)
}

func (g Grid[T]) ForEachCell(cellhandler cellhandler[T]) {
	for r := range g.Rows() {
		for c := range g.Columns() {
			loc := image.Point{r, c}
			if !cellhandler(loc, g.Cell(loc)) {
				break
			}
		}
	}
}

func (g Grid[T]) ForEachNeighbour(loc image.Point, cellhandler cellhandler[T]) {
	for _, d := range Directions4 {
		n := loc.Add(d.Point)
		if g.InGrid(n) {
			if !cellhandler(n, g.Cell(n)) {
				break
			}
		}
	}
}

func (g Grid[T]) NearBoundary(loc image.Point, dist int) bool {
	return loc.X <= dist || loc.Y <= dist ||
		loc.X >= g.Rows()-dist-1 || loc.Y >= g.Columns()-dist-1
}

func (g Grid[T]) InGrid(loc image.Point) bool {
	return loc.X >= 0 && loc.X < g.Rows() && loc.Y >= 0 && loc.Y < g.Columns()
}

func (g Grid[T]) FindFirst(target T) (image.Point, bool) {
	for r := range g.Rows() {
		for c := range g.Columns() {
			loc := image.Point{r, c}
			if g.Cell(loc) == target {
				return loc, true
			}
		}
	}
	return image.Point{0, 0}, false
}

func (g Grid[T]) WordAt(loc image.Point, length int, dir Direction) []T {
	word := []T{}
	limit := dir.Mul(length - 1)
	if !g.InGrid(loc.Add(limit)) {
		return word
	}

	for range length {
		word = append(word, g.Cell(loc))
		loc = loc.Add(dir.Point)
	}
	return word
}
