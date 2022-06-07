package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder

	var previous rune

	for i, r := range input {
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) {
			if previous == 0 {
				return "", ErrInvalidString
			}

			repeat, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(previous), repeat))

			previous = 0
		} else {
			if previous != 0 {
				result.WriteRune(previous)
			}
			previous = r
		}
	}

	if previous != 0 {
		result.WriteRune(previous)
	}

	return result.String(), nil
}
