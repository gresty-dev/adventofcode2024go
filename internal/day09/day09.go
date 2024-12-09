package day09

import (
	"bufio"
	"fmt"
	"io"

	. "go.gresty.dev/aoc2024/internal/lib"
)

type block struct {
	start, length, id int
}

func (b block) checksum() int64 {
	return int64(b.id) * int64(b.start*b.length+b.length*(b.length-1)/2)
}

func (b *block) trimleft(trim int) {
	b.start += trim
	b.length -= trim
}

type diskpointer struct {
	diskmap   []byte
	mapIndex  int
	fileBlock int
	block     int
}

func firstBlock(diskmap []byte) diskpointer {
	return diskpointer{diskmap: diskmap, mapIndex: 0, fileBlock: 0, block: 0}
}

func lastFileBlock(diskmap []byte) diskpointer {
	lastIndex := len(diskmap) - 1
	if lastIndex%2 == 1 { // it's the space after the last file
		lastIndex--
	}
	return diskpointer{diskmap: diskmap, mapIndex: lastIndex, fileBlock: int(diskmap[lastIndex] - 1), block: -1}
}

func (d *diskpointer) prevFileBlock() {
	d.fileBlock--
	for d.fileBlock == -1 {
		d.mapIndex -= 2
		d.fileBlock = d.blockSize() - 1
	}
}

func (d *diskpointer) nextBlock() {
	d.block++
	d.fileBlock++
	for d.fileBlock == d.blockSize() {
		d.fileBlock = 0
		d.mapIndex++
	}
}

func (d diskpointer) blockSize() int {
	return int(d.diskmap[d.mapIndex])
}

func (d diskpointer) isFile() bool {
	return d.mapIndex%2 == 0
}

func (d diskpointer) fileId() int {
	if d.isFile() {
		return d.mapIndex / 2
	} else {
		return -1
	}
}

func (d diskpointer) blockNumber() int {
	return d.block
}

func (d diskpointer) after(other diskpointer) bool {
	if d.mapIndex < other.mapIndex {
		return false
	}
	if d.mapIndex > other.mapIndex {
		return true
	}
	if d.fileBlock <= other.fileBlock {
		return false
	}
	return true
}

func Execute(input io.Reader) (Result, Result) {
	diskmap := readInput(input)
	r1 := NewResult(func() any { return defragByBlock(diskmap) })
	r2 := NewResult(func() any { return defragByFile(diskmap) })
	return r1, r2
}

func defragByBlock(diskmap []byte) int64 {
	p1 := firstBlock(diskmap)
	p2 := lastFileBlock(diskmap)

	done := false
	var checksum int64 = 0
	for !done {
		if p1.isFile() {
			checksum += int64(p1.blockNumber()) * int64(p1.fileId())
		} else {
			checksum += int64(p1.blockNumber()) * int64(p2.fileId())
			p2.prevFileBlock()
		}
		p1.nextBlock()
		done = p1.after(p2)
	}
	return checksum
}

func defragByFile(diskmap []byte) int64 {
	free, files := freeAndFiles(diskmap)
	searchFrom := [10]int{}
	checksum := int64(0)

	for f := len(files) - 1; f >= 0; f-- {
		moveto, found := search(searchFrom, free, files[f])
		if found {
			files[f].start = free[moveto].start
			free[moveto].trimleft(files[f].length)
		}
		checksum += files[f].checksum()
	}
	return checksum
}

func search(searchFrom [10]int, free []block, file block) (int, bool) {
	for i := searchFrom[file.length]; i < len(free); i++ {
		if free[i].start > file.start {
			break
		}
		if free[i].length >= file.length {
			searchFrom[file.length] = i
			return i, true
		}
	}
	return 0, false
}

func freeAndFiles(diskmap []byte) ([]block, []block) {
	free := make([]block, len(diskmap)/2+1)
	files := make([]block, len(diskmap)/2+1)
	s := 0
	for i, l := range diskmap {
		if i%2 == 1 {
			free[i/2] = block{start: s, length: int(l)}
		} else {
			files[i/2] = block{start: s, length: int(l), id: i / 2}
		}
		s += int(l)
	}
	return free, files
}

func readInput(input io.Reader) []byte {
	diskmap := []byte{}
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	for _, b := range scanner.Bytes() {
		diskmap = append(diskmap, b-'0')
	}
	fmt.Println("Read ", len(diskmap), " entries")
	return diskmap
}
