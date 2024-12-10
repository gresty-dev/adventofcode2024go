package day06

import (
	"image"
	"io"

	aoc "go.gresty.dev/aoc2024/internal/lib"
)

type void struct{}

var empty void

type Guard struct {
	image.Point
	*aoc.Direction
}

func (g *Guard) turnRight() {
	g.Direction = g.Direction.Next90()
}

func (g *Guard) move() {
	g.Point = g.Point.Add(g.Direction.Point)
}

type visitedCell struct {
	image.Point
	aoc.Direction
}

func Execute(input io.Reader) (aoc.Result, aoc.Result) {
	grid := aoc.ReadGrid(input)

	var guard Guard
	var visited map[image.Point]void

	result1 := aoc.NewResult(func() any {
		guardLoc, ok := grid.FindFirst('^')
		if !ok {
			panic("no guard found!")
		}
		guard = Guard{guardLoc, &aoc.Up}
		visited = walk(guard, grid)
		return len(visited)
	})

	result2 := aoc.NewResult(func() any {
		return loops(guard, grid, visited)
	})

	return result1, result2
}

func walk(guard Guard, grid aoc.Grid[byte]) map[image.Point]void {
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

func loops(guard Guard, grid aoc.Grid[byte], obstacles map[image.Point]void) int {
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

func walkWithObstacle(guard Guard, grid aoc.Grid[byte], obstacle image.Point) bool {
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
