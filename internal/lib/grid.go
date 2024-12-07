package lib

import (
	"bufio"
	"image"
	"io"
)

type cellhandler func(loc image.Point, value byte) bool

type Grid struct {
	cells [][]byte
}

func ReadGrid(input io.Reader) Grid {
	grid := [][]byte{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}
	return Grid{grid}
}

func (g Grid) Cell(loc image.Point) byte {
	return g.cells[loc.X][loc.Y]
}

func (g *Grid) Set(loc image.Point, value byte) {
	g.cells[loc.X][loc.Y] = value
}

func (g Grid) Columns() int {
	return len(g.cells[0])
}

func (g Grid) Rows() int {
	return len(g.cells)
}

func (g Grid) ForEachCell(cellhandler cellhandler) {
	for r := range g.Rows() {
		for c := range g.Columns() {
			loc := image.Point{r, c}
			if !cellhandler(loc, g.Cell(loc)) {
				break
			}
		}
	}
}

func (g Grid) NearBoundary(loc image.Point, dist int) bool {
	return loc.X <= dist || loc.Y <= dist ||
		loc.X >= g.Rows()-dist-1 || loc.Y >= g.Columns()-dist-1
}

func (g Grid) InGrid(loc image.Point) bool {
	return loc.X >= 0 && loc.X < g.Rows() && loc.Y >= 0 && loc.Y < g.Columns()
}

func (g Grid) FindFirst(target byte) (image.Point, bool) {
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

func (g Grid) WordAt(loc image.Point, length int, dir Direction) string {
	limit := dir.Mul(length - 1)
	if !g.InGrid(loc.Add(limit)) {
		return ""
	}

	word := []byte{}
	for range length {
		word = append(word, g.Cell(loc))
		loc = loc.Add(dir.Point)
	}
	return string(word)
}
