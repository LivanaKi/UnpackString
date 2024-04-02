package unpackstring

import (
	"strconv"
	"strings"
	"unicode"
	"errors"
)

var ErrInvalidString = errors.New("Number it isn't ok")

func Unpack(s string) (string, error){

	var res strings.Builder
	
	if len(s) == 0{
		return "", nil
	}

	if unicode.IsNumber(rune(s[0])){
		return "",ErrInvalidString
	}else{
		res.WriteString(string(s[0]))
	}

	for i := 1; i < len(s); i++{
		if unicode.IsLetter(rune(s[i])){
			res.WriteString(string(s[i]))
		}else{
			var current string
			num, _ := strconv.Atoi(string(s[i]))
			if num == 0 {
				return "", ErrInvalidString
			}
			current = strings.Repeat(string(s[i-1]), num - 1)

			res.WriteString(current)
		}
	}
	return res.String(), nil
}