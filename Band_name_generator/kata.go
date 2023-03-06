package kata

import "strings"

func bandNameGenerator(word string) string {
	var res string

	if word[0] == word[len(word)-1] {
		res = strings.Title(word) + word[1:]
	} else {
		res = "The " + strings.Title(word)
	}

	return res
}
