package kata

func GetMiddle(s string) string {
	sLen := len(s)

	if sLen == 0 || sLen == 1 {
		return s
	}

	var half int
	if sLen%2 == 0 {
		half = (sLen / 2) - 1
	} else {
		half = (sLen / 2)
	}

	res := s[half : sLen-half]
	return res
}
