package kata

import (
	"strings"
)

func Capitalize(st string) []string {
	var evenString, oddString string

	for i, ch := range st {
		letter := string(ch)
		if i%2 == 0 {
			evenString += strings.ToUpper(letter)
			oddString += strings.ToLower(letter)
		} else {
			evenString += strings.ToLower(letter)
			oddString += strings.ToUpper(letter)
		}
	}
	return []string{evenString, oddString}
}
