package main

import (
	"fmt"
	"piscine"
)

func main() {
	x, y, flag := piscine.GetNumber()
	if !flag {
		return
	}
	fmt.Print(piscine.QuadB(x, y))
}
