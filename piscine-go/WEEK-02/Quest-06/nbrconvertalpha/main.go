package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	charCase := 0
	if args[0] == "--upper" {
		charCase = -32
		args = args[1:]
	}

	for _, a := range args {
		digit := 0
		for _, b := range a {
			digit = digit*10 + int(b-'0')
		}
		if 1 <= digit && digit <= 26 {
			z01.PrintRune('a' + rune(digit+charCase) - 1)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
