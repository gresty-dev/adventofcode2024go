package lib

import "image"

type Direction struct {
	image.Point
	prev, next *Direction
}

func (d Direction) Next() *Direction {
	return d.next
}

func (d Direction) Next90() *Direction {
	return d.next.next
}

func (d Direction) Prev() *Direction {
	return d.prev
}

func (d Direction) Prev90() *Direction {
	return d.prev.prev
}

// X maps to row, Y maps to column!!
var Directions = []Direction{
	Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft,
}

var Directions4 = []Direction{
	Up, Right, Down, Left,
}

var Up = Direction{image.Point{-1, 0}, nil, nil}
var UpRight = Direction{image.Point{-1, 1}, &Up, nil}
var Right = Direction{image.Point{0, 1}, &UpRight, nil}
var DownRight = Direction{image.Point{1, 1}, &Right, nil}
var Down = Direction{image.Point{1, 0}, &DownRight, nil}
var DownLeft = Direction{image.Point{1, -1}, &Down, nil}
var Left = Direction{image.Point{0, -1}, &DownLeft, nil}
var UpLeft = Direction{image.Point{-1, -1}, &Left, &Up}

func init() {
	Up.prev = &UpLeft
	Up.next = &UpRight
	UpRight.next = &Right
	Right.next = &DownRight
	DownRight.next = &Down
	Down.next = &DownLeft
	DownLeft.next = &Left
	Left.next = &UpLeft
}
