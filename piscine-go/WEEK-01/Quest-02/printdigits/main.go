package main

import "github.com/01-edu/z01"

func main() {
	for digit := '0'; digit <= '9'; digit++ {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}
