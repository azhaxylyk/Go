package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	count := 0
	for tens1 := 9; tens1 >= 0; tens1-- {
		for ones1 := 9; ones1 >= 0; ones1-- {
			for tens2 := 9; tens2 >= 0; tens2-- {
				for ones2 := 9; ones2 >= 0; ones2-- {
					if tens1*10+ones1 > tens2*10+ones2 {
						if count != 0 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune('0' + tens1))
						z01.PrintRune(rune('0' + ones1))
						z01.PrintRune(' ')
						z01.PrintRune(rune('0' + tens2))
						z01.PrintRune(rune('0' + ones2))
						count++
					}
				}
			}
		}
	}
}
