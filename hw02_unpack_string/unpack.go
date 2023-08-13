package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sr string) (string, error) {

	var last_char strings.Builder
	var full_string strings.Builder
	var last_char_is_digit bool

	for pos, char := range sr {

		if unicode.IsDigit(char) && pos == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(char) && len(last_char.String()) == 0 {
			last_char.WriteString(string(char))
			full_string.WriteString(string(char))
			last_char_is_digit = false
			continue
		}

		if unicode.IsLetter(char) && len(last_char.String()) > 0 {
			last_char.Reset()
			last_char.WriteString(string(char))
			full_string.WriteString(string(char))
			last_char_is_digit = false
			continue
		}

		if unicode.IsDigit(char) && last_char_is_digit == true {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && len(last_char.String()) > 0 {
			y, _ := strconv.Atoi(string(char))

			if y == 0 {
				str := full_string.String()
				str = str[:len(str)-1]
				full_string.Reset()
				full_string.WriteString(str)
			}

			if y > 0 {
				for i := 1; i < y; i++ {
					full_string.WriteString(last_char.String())
				}
			}
			last_char.Reset()
			last_char_is_digit = true
			continue
		}

	}

	return full_string.String(), nil
}
