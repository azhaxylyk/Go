package piscine

func ActiveBits(n int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			count++
		}
	}
	return count
}
