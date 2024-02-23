package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n > 0 && n <= len(r) {
		return r[n-1]
	} else {
		return 0
	}
}
