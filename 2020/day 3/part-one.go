package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	EMPTY = '.'
	TREE  = '#'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	grid := readInput(file)

	treeCount := countTreesInWalk(grid, 3, 1)

	fmt.Println(treeCount)
}

func readInput(r io.Reader) [][]rune {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var lines []string
	row := 0
	for s.Scan() {
		lines = append(lines, s.Text())
		row++
	}

	var result = make2dSlice(len(lines[0]), row)

	for i, v := range lines {
		for j, r := range v {
			result[i][j] = r
		}
	}

	return result
}

func make2dSlice(w int, h int) [][]rune {
	a := make([]rune, w*h)
	s := make([][]rune, h)
	lo, hi := 0, w
	for i := range s {
		s[i] = a[lo:hi:hi]
		lo, hi = hi, hi+w
	}
	return s
}

func countTreesInWalk(grid [][]rune, x_step int, y_step int) int {
	treeCount := 0

	px, py := 0, 0
	for py < len(grid) {
		// check current position for tree
		if grid[py][px] == TREE {
			treeCount++
		}

		// move our tree along
		px += x_step
		for px >= len(grid[py]) {
			px = px - len(grid[py])
		}
		py += y_step
	}

	return treeCount
}
