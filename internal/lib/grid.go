package lib

import (
	"bufio"
	"fmt"
	"image"
	"io"
)

type cellhandler[T comparable] func(loc image.Point, value T)

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
			cellhandler(loc, g.Cell(loc))
		}
	}
}

func (g Grid[T]) ForEachNeighbour(loc image.Point, cellhandler cellhandler[T]) {
	for _, d := range Directions4 {
		n := loc.Add(d.Point)
		if g.InGrid(n) {
			cellhandler(n, g.Cell(n))
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

func (g *Grid[T]) MergeNonZero(origin image.Point, other Grid[T]) {
	var zero T
	other.ForEachCell(func(loc image.Point, value T) {
		if value != zero {
			g.Set(origin.Add(loc), value)
		}
	})
}

func (g Grid[T]) FindArea(loc image.Point) Area[T] {
	area := Area[T]{
		match:  g.Cell(loc),
		Grid:   NewGrid[bool](g.Rows(), g.Columns()),
		Bounds: rectFromPoint(loc),
	}
	g.walkArea(loc, &area)
	area.crop()
	return area
}

func (g *Grid[T]) crop(rect image.Rectangle) {
	g.cells = g.cells[rect.Min.X:rect.Max.X]
	for r := 0; r < rect.Dx(); r++ {
		g.cells[r] = g.cells[r][rect.Min.Y:rect.Max.Y]
	}
}

func (g Grid[T]) walkArea(loc image.Point, area *Area[T]) {
	area.include(loc)
	g.ForEachNeighbour(loc, func(n image.Point, value T) {
		if area.match == value && !area.Grid.Cell(n) {
			g.walkArea(n, area)
		}
	})
}

func (g Grid[T]) Dump() {
	for r := range g.Rows() {
		for c := range g.Columns() {
			fmt.Print(g.cells[r][c])
		}
		fmt.Println()
	}
}

type Area[T comparable] struct {
	match  T
	Grid   Grid[bool]
	Bounds image.Rectangle
	Origin image.Point
	Size   int
}

func (a *Area[T]) include(loc image.Point) {
	if !a.Grid.Cell(loc) {
		a.Grid.Set(loc, true)
		a.Size++
	}
	a.Bounds = a.Bounds.Union(rectFromPoint(loc))
}

func (a Area[T]) Perimeter() int {
	perimeter := 0
	a.Grid.ForEachCell(func(loc image.Point, inArea bool) {
		if inArea {
			if loc.X == 0 {
				perimeter++
			}
			if loc.Y == 0 {
				perimeter++
			}
			if loc.X == a.Grid.Rows()-1 {
				perimeter++
			}
			if loc.Y == a.Grid.Columns()-1 {
				perimeter++
			}
			a.Grid.ForEachNeighbour(loc, func(n image.Point, nInArea bool) {
				if !nInArea {
					perimeter++
				}
			})
		}
	})
	return perimeter
}

func (a *Area[T]) crop() {
	a.Grid.crop(a.Bounds.Add(a.Origin))
	a.Origin = a.Origin.Add(a.Bounds.Min)
	a.Bounds = a.Bounds.Sub(a.Bounds.Min)
}

func rectFromPoint(p image.Point) image.Rectangle {
	return image.Rectangle{p, p.Add(image.Point{1, 1})}
}
