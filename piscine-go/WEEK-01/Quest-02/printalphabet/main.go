package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, ch := range alphabet {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
