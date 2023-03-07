package kata

func RoundToNext5(n int) int {
	if n != 0 {
		for n%5 != 0 {
			n++
		}
		return n
	}
	return 0
}
