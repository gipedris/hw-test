package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sr string) (string, error) {

	var last_char strings.Builder
	var full_string strings.Builder

	for pos, char := range sr {

		if pos == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(char) && len(last_char.String()) == 0 {
			last_char.WriteString(string(char))
			full_string.WriteString(string(char))
			continue
		}

		if unicode.IsLetter(char) && len(last_char.String()) > 0 {
			last_char.Reset()
			last_char.WriteString(string(char))
			full_string.WriteString(string(char))
			continue
		}

		if unicode.IsDigit(char) {
			y, _ := strconv.Atoi(string(char))

			for i := 1; i < y; i++ {
				full_string.WriteString(last_char.String())
			}
			last_char.Reset()
			continue
		}

	}

	return full_string.String(), nil
}

func main() {
	a, err := Unpack("ab3a3c6")
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
	}
}
