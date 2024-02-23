package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for !IsPrime(n) {
			n++
		}
		return n
	}
}
