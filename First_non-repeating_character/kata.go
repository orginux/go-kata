package kata

import "strings"

func FirstNonRepeating(str string) string {
	characters := make(map[string]int)

	for _, ch := range str {
		letter := strings.ToLower(string(ch))
		characters[letter]++
	}

	for _, ch := range str {
		letter := strings.ToLower(string(ch))
		if characters[letter] == 1 {
			return string(ch)
		}
	}
	return ""
}
