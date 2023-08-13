package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sr string) (string, error) {
	var lastChar strings.Builder
	var fullString strings.Builder
	var lastCharIsDigit bool

	for pos, char := range sr {
		if unicode.IsDigit(char) && pos == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(char) && len(lastChar.String()) == 0 {
			lastChar.WriteString(string(char))
			fullString.WriteString(string(char))
			lastCharIsDigit = false
			continue
		}

		if unicode.IsLetter(char) && len(lastChar.String()) > 0 {
			lastChar.Reset()
			lastChar.WriteString(string(char))
			fullString.WriteString(string(char))
			lastCharIsDigit = false
			continue
		}

		if unicode.IsDigit(char) && lastCharIsDigit {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && len(lastChar.String()) > 0 {
			y, _ := strconv.Atoi(string(char))

			if y == 0 {
				str := fullString.String()
				str = str[:len(str)-1]
				fullString.Reset()
				fullString.WriteString(str)
			}

			if y > 0 {
				for i := 1; i < y; i++ {
					fullString.WriteString(lastChar.String())
				}
			}
			lastChar.Reset()
			lastCharIsDigit = true
			continue
		}
	}

	return fullString.String(), nil
}
