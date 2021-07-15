package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	min       int
	max       int
	character rune
	password  string
}

func (pp *PasswordPolicy) isValid() bool {
	return (pp.password[pp.min-1] == byte(pp.character) && pp.password[pp.max-1] != byte(pp.character)) ||
		(pp.password[pp.max-1] == byte(pp.character) && pp.password[pp.min-1] != byte(pp.character))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	input, err := ReadInput(file)

	output := countValidPasswords(input)

	fmt.Println(output)
}

func ReadInput(r io.Reader) ([]PasswordPolicy, error) {
	var result []PasswordPolicy

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		ints := strings.Split(line[0], "-")

		min, err := strconv.Atoi(ints[0])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(ints[1])
		if err != nil {
			return nil, err
		}

		pp := PasswordPolicy{
			min:       min,
			max:       max,
			character: rune(line[1][0]),
			password:  line[2],
		}
		result = append(result, pp)
	}
	return result, s.Err()
}

func countValidPasswords(passwords []PasswordPolicy) int {
	count := 0
	for _, p := range passwords {
		if p.isValid() {
			count++
		}
	}

	return count
}
