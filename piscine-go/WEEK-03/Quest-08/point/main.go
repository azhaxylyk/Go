package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point, x, y int) {
	ptr.x = x
	ptr.y = y
}

func printInt(xy int) {
	var digits []rune
	for xy > 0 {
		digit := rune('0' + (xy % 10))
		digits = append([]rune{digit}, digits...)
		xy /= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func main() {
	points := point{}

	setPoint(&points, 42, 21)
	xmessage := "x = "
	ymessage := ", y = "
	for _, i := range xmessage {
		z01.PrintRune(i)
	}
	printInt(points.x)
	for _, j := range ymessage {
		z01.PrintRune(j)
	}
	printInt(points.y)
	z01.PrintRune('\n')
}
