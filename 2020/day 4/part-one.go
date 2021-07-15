package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	BIRTH_YEAR      = "byr"
	ISSUE_YEAR      = "iyr"
	EXPIRATION_YEAR = "eyr"
	HEIGHT          = "hgt"
	HAIR_COLOR      = "hcl"
	EYE_COLOR       = "ecl"
	PASSPORT_ID     = "pid"
	COUNTRY_ID      = "cid"
)

var IgnoredFields = make(map[string]struct{})
var AMOUNT_OF_UNIQUE_VALID_FIELDS_REQUIRED = 7

type Passport struct {
	ValidFields   map[string]string
	IgnoredFields map[string]string
}

func (passport Passport) isValid() bool {
	return len(passport.ValidFields) >= AMOUNT_OF_UNIQUE_VALID_FIELDS_REQUIRED
}

func NewPassport() Passport {
	return Passport{
		ValidFields:   make(map[string]string),
		IgnoredFields: make(map[string]string),
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	IgnoredFields[COUNTRY_ID] = struct{}{}
	passports := readInput(file)

	validPassportCount := 0
	for _, p := range passports {
		if p.isValid() {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}

func readInput(r io.Reader) []Passport {
	var passports []Passport

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	p := NewPassport()
	for s.Scan() {
		line := s.Text()

		if line == "" {
			passports = append(passports, p)
			p = NewPassport()
			continue
		}

		attributes := strings.Split(line, " ")
		for _, a := range attributes {
			s := strings.Split(a, ":")
			_, ignored := IgnoredFields[s[0]]
			if !ignored {
				p.ValidFields[s[0]] = s[1]
			} else {
				p.IgnoredFields[s[0]] = s[1]
			}
		}
	}
	// append the final passport
	passports = append(passports, p)

	return passports
}
