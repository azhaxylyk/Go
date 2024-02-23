package piscine

func IterativePower(nb int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	if power > 0 {
		return result
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
