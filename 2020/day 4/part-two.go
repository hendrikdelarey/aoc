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
	InvalidFields map[string]string
}

func (passport Passport) isValid() bool {
	return len(passport.ValidFields) >= AMOUNT_OF_UNIQUE_VALID_FIELDS_REQUIRED
}

func NewPassport() Passport {
	return Passport{
		ValidFields:   make(map[string]string),
		IgnoredFields: make(map[string]string),
		InvalidFields: make(map[string]string),
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
			if ignored {
				p.IgnoredFields[s[0]] = s[1]
			} else if isValidPassportField(s[0], s[1]) {
				p.ValidFields[s[0]] = s[1]
			} else {
				p.InvalidFields[s[0]] = s[1]
			}
		}
	}
	// append the final passport
	passports = append(passports, p)

	return passports
}

func isValidPassportField(key string, value string) bool {
	switch key {
	case BIRTH_YEAR:
		return isValidBirthYear(value)
	case ISSUE_YEAR:
		return isValidIssueYear(value)
	case EXPIRATION_YEAR:
		return isValidExpirationYear(value)
	case HEIGHT:
		return isValidHeight(value)
	case HAIR_COLOR:
		return isValidHairColor(value)
	case EYE_COLOR:
		return isValidEyeColor(value)
	case PASSPORT_ID:
		return isValidPassportId(value)
	}

	return false
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func isValidBirthYear(year string) bool {
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return y >= 1920 && y <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func isValidIssueYear(year string) bool {
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return y >= 2010 && y <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func isValidExpirationYear(year string) bool {
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return y >= 2020 && y <= 2030
}

/*
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
*/
func isValidHeight(height string) bool {
	if len(height) < 4 {
		return false
	}
	unit := height[len(height)-2:]

	n, err := strconv.Atoi(height[0 : len(height)-2])
	if err != nil {
		return false
	}

	if unit == "cm" {
		return n >= 150 && n <= 193
	} else if unit == "in" {
		return n >= 59 && n <= 76
	}

	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isValidHairColor(color string) bool {
	if len(color) != 7 {
		return false
	}

	if color[0] != '#' {
		return false
	}

	return isValidHexDigit(color[1:])
}

var validEyeColors = map[string]struct{}{
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {}}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func isValidEyeColor(color string) bool {
	_, iv := validEyeColors[color]
	return iv
}

var validPassportDigits = map[rune]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {}}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isValidPassportId(passportId string) bool {
	if len(passportId) != 9 {
		return false
	}

	for _, d := range passportId {
		_, vd := validPassportDigits[d]
		if !vd {
			return false
		}
	}

	return true
}

func isValidHexDigit(val string) bool {
	_, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		return false
	}
	return true
}
