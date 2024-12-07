package day06

import (
	"image"
	"io"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

type void struct{}

var empty void

type guard struct {
	image.Point
	*aoc.Direction
}

func (g *guard) turnRight() {
	g.Direction = g.Direction.Next90()
}

func (g *guard) move() {
	g.Point = g.Point.Add(g.Direction.Point)
}

type visitedCell struct {
	image.Point
	aoc.Direction
}

func Execute(input io.Reader) (any, any) {
	grid := aoc.ReadGrid(input)
	guardLoc, ok := grid.FindFirst('^')
	if !ok {
		panic("no guard found!")
	}
	guard := guard{guardLoc, &aoc.Up}

	visited := walk(guard, grid)

	return len(visited), loops(guard, grid, visited)
}

func walk(guard guard, grid aoc.Grid) map[image.Point]void {
	visited := make(map[image.Point]void)
	for grid.InGrid(guard.Point) {
		visited[guard.Point] = empty
		next := guard.Point.Add(guard.Direction.Point)
		if grid.InGrid(next) && grid.Cell(next) == '#' {
			guard.turnRight()
		} else {
			guard.move()
		}
	}
	return visited
}

func loops(guard guard, grid aoc.Grid, obstacles map[image.Point]void) int {
	count := 0
	for obstacle := range obstacles {
		if obstacle == guard.Point {
			continue
		}
		if walkWithObstacle(guard, grid, obstacle) {
			count++
		}
	}
	return count
}

func walkWithObstacle(guard guard, grid aoc.Grid, obstacle image.Point) bool {
	visited := make(map[visitedCell]void)
	grid.Set(obstacle, '#')
	defer grid.Set(obstacle, '.')

	for grid.InGrid(guard.Point) {
		currentCell := visitedCell{guard.Point, *guard.Direction}
		_, ok := visited[currentCell]
		if ok {
			return true
		}
		next := guard.Point.Add(guard.Direction.Point)
		if grid.InGrid(next) && grid.Cell(next) == '#' {
			visited[currentCell] = empty // optimisation - only track turns, not all cells on path
			guard.turnRight()
		} else {
			guard.move()
		}
	}
	return false
}