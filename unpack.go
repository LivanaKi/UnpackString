package unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(start string) (string, error) {

	var result strings.Builder

	if len(start) == 0 {
		return "", nil
	}
	for i, solution := range start {
		if unicode.IsNumber(rune(start[0])) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(solution) {
			num, _ := strconv.Atoi(string(start[i]))
			if num == 1 && unicode.IsDigit(rune(start[i+1])) {
				num1, _ := strconv.Atoi(string(start[i+1]))
				if num1 == 0 {
					return "", ErrInvalidString
				}
			}
			if num == 0 {
				return "", ErrInvalidString
			}
			result.WriteString(strings.Repeat(string(start[i-1]), num-1))
		}
		if unicode.IsLetter(solution) {
			result.WriteRune(solution)
		}
	}
	return result.String(), nil
}
