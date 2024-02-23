package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	sqrt := 1
	for sqrt*sqrt <= nb {
		sqrt++
	}
	sqrt--

	for i := 3; i <= sqrt; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
