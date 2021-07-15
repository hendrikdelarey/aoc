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

	var groups []map[rune]struct{}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	group := make(map[rune]struct{})
	for s.Scan() {
		line := s.Text()

		if line == "" {
			groups = append(groups, group)
			group = make(map[rune]struct{})
		}

		for _, v := range line {
			group[v] = struct{}{}
		}
	}
	groups = append(groups, group)

	voteCount := 0
	for _, g := range groups {
		voteCount += len(g)
	}

	fmt.Printf("Vote Count: %d\n", voteCount)
}
