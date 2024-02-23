package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]

	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			programName = programName[i+1:]
			break
		}
	}

	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
