package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	group := make(map[rune]int)
	groupSize := 0
	voteCount := 0
	for s.Scan() {
		line := s.Text()

		// group is done voting
		if line == "" {
			// for all the votes in the group
			for _, gv := range group {
				if gv == groupSize {
					voteCount++
				}
			}
			group = make(map[rune]int)
			groupSize = 0
			continue
		}

		// add each individual vote to count
		for _, v := range line {
			cv, ok := group[v]
			if !ok {
				cv = 0
			}
			group[v] = cv + 1
		}

		// count the amount of voters in the group
		groupSize++
	}
	// don't forget about the last group
	for _, gv := range group {
		if gv == groupSize {
			voteCount++
		}
	}

	fmt.Printf("Vote Count: %d\n", voteCount)
}
