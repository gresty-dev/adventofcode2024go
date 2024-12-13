package day12

import (
	"image"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type region struct {
	plots              Grid[bool]
	origin             image.Point
	area, fence, sides int
}

func Execute(input io.Reader) (Result, Result) {
	field := ReadGrid(input)
	var p1, p2 int

	r1 := NewResult(func() any {
		p1, p2 = fencingPrices(field)
		return p1
	})

	r2 := NewResult(func() any {
		return p2
	})

	return r1, r2
}

func fencingPrices(field Grid[byte]) (int, int) {
	price1 := 0
	price2 := 0
	mapped := NewGrid[bool](field.Rows(), field.Columns())
	field.ForEachCell(func(loc image.Point, value byte) {
		if !mapped.Cell(loc) {
			region := mapRegion(field, loc)
			price1 += region.area * region.fence
			price2 += region.area * region.sides
			mapped.MergeNonZero(region.origin, region.plots)
		}
	})
	return price1, price2
}

func mapRegion(field Grid[byte], loc image.Point) region {
	area := field.FindArea(loc)
	return region{
		plots:  area.Grid,
		origin: area.Origin,
		area:   area.Size,
		fence:  area.Perimeter(),
		sides:  countSides(area),
	}
}

func countSides(area Area[byte]) int {
	corners := map[image.Point]Void{}
	diagonals := map[image.Point]Void{}
	area.Grid.ForEachCell(func(loc image.Point, value bool) {
		inArea := area.Grid.Cell(loc)
		up := neighbourInArea(area.Grid, loc, Up)
		down := neighbourInArea(area.Grid, loc, Down)
		left := neighbourInArea(area.Grid, loc, Left)
		right := neighbourInArea(area.Grid, loc, Right)
		upright := neighbourInArea(area.Grid, loc, UpRight)
		downright := neighbourInArea(area.Grid, loc, DownRight)

		if up == !inArea && left == !inArea {
			corners[loc] = Empty
		}
		if up == !inArea && right == !inArea {
			corners[loc.Add(Right.Point)] = Empty
		}
		if down == !inArea && left == !inArea {
			corners[loc.Add(Down.Point)] = Empty
		}
		if down == !inArea && right == !inArea {
			corners[loc.Add(Down.Point).Add(Right.Point)] = Empty
		}

		if inArea && downright && !right && !down {
			diagonals[loc.Add(Down.Point).Add(Right.Point)] = Empty
		}
		if inArea && upright && !right && !up {
			diagonals[loc.Add(Up.Point).Add(Right.Point)] = Empty
		}
	})
	return len(corners) + len(diagonals)
}

func neighbourInArea(grid Grid[bool], loc image.Point, dir Direction) bool {
	neighbour := loc.Add(dir.Point)
	if grid.InGrid(neighbour) {
		return grid.Cell(neighbour)
	}
	return false
}
