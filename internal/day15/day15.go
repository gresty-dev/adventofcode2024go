package day15

import (
	"bufio"
	"image"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type Warehouse struct {
	Grid[byte]
	robot image.Point
}

func (w *Warehouse) moveRobot(d Direction) {
	next := w.robot.Add(d.Point)
	if w.Cell(next) == '#' {
		return
	}
	if w.Cell(next) == '.' {
		w.robot = next
		return
	}
	if w.Cell(next) == 'O' && w.moveBox(next, d) {
		w.robot = next
		return
	}
}

func (w *Warehouse) moveBox(box image.Point, d Direction) bool {
	next := box.Add(d.Point)
	if w.Cell(next) == '#' {
		return false
	}
	if w.Cell(next) == '.' || w.moveBox(next, d) {
		w.Set(next, 'O')
		w.Set(box, '.')
		return true
	}
	return false
}

func (w Warehouse) sumBoxCoords() int {
	sum := 0
	w.ForEachCell(func(cell image.Point, val byte) {
		if val == 'O' {
			sum += 100*cell.X + cell.Y
		}
	})
	return sum
}

func Execute(input io.Reader) (Result, Result) {
	warehouse, instr := readInput(input)

	r1 := NewResult(func() any {
		return runSim(warehouse, instr)
	})
	r2 := NewResult(func() any {
		return 0
	})
	return r1, r2
}

func runSim(w Warehouse, instr []byte) int {
	for _, i := range instr {
		w.moveRobot(direction(i))
	}
	return w.sumBoxCoords()
}

func direction(instr byte) Direction {
	switch instr {
	case '^':
		return Up
	case '>':
		return Right
	case 'v':
		return Down
	case '<':
		return Left
	}
	return Direction{Point: image.Point{0, 0}}
}

func readInput(input io.Reader) (Warehouse, []byte) {
	w := Warehouse{}

	scanner := bufio.NewScanner(input)
	w.Grid = ScanGrid(scanner)
	var ok bool
	if w.robot, ok = w.FindFirst('@'); !ok {
		panic("No robot found")
	}
	w.Set(w.robot, '.')

	instr := []byte{}
	for scanner.Scan() {
		line := scanner.Text()
		instr = append(instr, line...)
	}

	return w, instr
}
