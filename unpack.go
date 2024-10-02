package unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	for id, val := range str {
		if unicode.IsNumber(rune(str[0])) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(val) {
			num, err := strconv.Atoi(string(str[id]))
			if err != nil {
				return "", err
			}
			if num >= 0 && unicode.IsDigit(rune(str[id+1])) || unicode.IsDigit(rune(str[len(str) - 1])){
				return "", ErrInvalidString
			}

			if num == 0 {
				str1 := result.String()
				str1 = str1[:len(str1)-1]
				result.Reset()
				result.WriteString(str1)

			} else {
				result.WriteString(strings.Repeat(string(str[id-1]), num-1))
			}
		}
		if unicode.IsLetter(val) {
			result.WriteRune(val)
		}
	}
	return result.String(), nil
}
