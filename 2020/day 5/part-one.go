package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	B = 'B'
	F = 'F'
	L = 'L'
	R = 'R'
)

const ROWS, COLS = 128, 8

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	highestSeatId := 0

	for s.Scan() {
		inst := s.Text()
		row := getRow(inst)
		col := getCol(inst)
		seat := getSeatID(row, col)

		if seat > highestSeatId {
			highestSeatId = seat
		}
	}

	fmt.Printf("Highest seat id: %d\n", highestSeatId)
}

func getRow(inst string) int {
	minRow, maxRow := 0, ROWS
	for _, i := range inst {
		if minRow == maxRow {
			return minRow
		}

		if i == B {
			minRow += (maxRow - minRow) / 2
		} else if i == F {
			maxRow -= (maxRow - minRow) / 2
		} else {
			break
		}
	}

	return minRow
}

func getCol(inst string) int {
	minCol, maxCol := 0, COLS
	for _, i := range inst {
		if minCol == maxCol {
			return minCol
		}

		if i == R {
			minCol += (maxCol - minCol) / 2
		} else if i == L {
			maxCol -= (maxCol - minCol) / 2
		}
	}

	return minCol
}

func getSeatID(row int, col int) int {
	return row*8 + col
}
