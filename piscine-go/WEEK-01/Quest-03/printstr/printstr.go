package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	Byte := []byte(s)

	for _, letter := range Byte {
		z01.PrintRune(rune(letter))
	}
}
