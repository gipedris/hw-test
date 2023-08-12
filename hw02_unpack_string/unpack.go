package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	sr := []rune(s)
	var current_char strings.Builder
	var full_string strings.Builder
	// var current_pos int

	for pos, char := range sr {

		if pos == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(char) {
			current_char.WriteRune(char)
			continue
		}

		if unicode.IsDigit(char) {
			y, _ := strconv.Atoi(string(char))

			for i := 0; i < y; i++ {
				full_string.WriteString(current_char.String())
				full_string.WriteString(current_char.String())
			}
			current_char.Reset()
			continue
		}

	}

	return full_string.String(), nil
}

func main() {
	a, err := Unpack("ab")
	fmt.Print(a)
	if err != nil {
		fmt.Print(err)
	}
}
