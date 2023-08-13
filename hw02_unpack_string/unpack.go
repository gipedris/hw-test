package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var lastChar rune
	var fullString strings.Builder
	var lastCharIsDigit bool

	for _, char := range s {
		if unicode.IsDigit(char) && lastChar == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(char) && lastChar == 0 {
			lastChar = char
			fullString.WriteRune(char)
			lastCharIsDigit = false
			continue
		}

		if unicode.IsLetter(char) && lastChar != 0 {
			lastChar = char
			fullString.WriteRune(char)
			lastCharIsDigit = false
			continue
		}

		if unicode.IsDigit(char) && lastCharIsDigit {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && !lastCharIsDigit {
			y, _ := strconv.Atoi(string(char))

			if y == 0 {
				a := strings.TrimSuffix(fullString.String(), string(lastChar))
				fullString.Reset()
				fullString.WriteString(a)
			}

			if y > 0 {
				for i := 1; i < y; i++ {
					fullString.WriteRune(lastChar)
				}
			}
			lastChar = char
			lastCharIsDigit = true
			continue
		}
	}

	return fullString.String(), nil
}
