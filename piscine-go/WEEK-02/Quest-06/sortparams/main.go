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

	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if compareStrings(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func compareStrings(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return len(a) - len(b)
}
