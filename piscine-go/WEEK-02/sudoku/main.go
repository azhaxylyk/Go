package main

import (
	"os"

	"github.com/01-edu/z01"
)

func solveSudoku(grid [][]int) bool {
	row, col := -1, -1
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				row, col = i, j
				break
			}
		}
		if row != -1 {
			break
		}
	}
	if row == -1 {
		return true
	}
	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func isValid(grid [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}
	subgridRow := row - row%3
	subgridCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[subgridRow+i][subgridCol+j] == num {
				return false
			}
		}
	}
	return true
}

func printGrid(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(grid[i][j] + '0'))
			if j != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}

func hasUniqueElements(strs string) bool {
	for _, ch := range strs {
		if ch != '.' && (ch < '0' || ch > '9') {
			return false
		}
	}
	for i := 0; i < 8; i++ {
		for j := 1; j < 9; j++ {
			if i != j && strs[i] == strs[j] && strs[i] != '.' {
				return false
			}
		}
	}
	return true
}

func stringToRuneAndPrint(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func hasUniqueSolution(table [][]int) bool {
	solutions := 0
	var backtrack func(int, int)
	ch := make(chan int, 2)
	backtrack = func(row, col int) {
		if len(ch) > 1 {
			return
		}
		if row == 9 {
			solutions++
			ch <- -1
			return
		}
		if col == 9 {
			backtrack(row+1, 0)
			return
		}
		if table[row][col] == 0 {
			for num := 1; num <= 9; num++ {
				if isValid(table, row, col, num) {
					table[row][col] = num
					backtrack(row, col+1)
					table[row][col] = 0
				}
			}
		} else {
			backtrack(row, col+1)
		}
	}
	backtrack(0, 0)
	return solutions == 1
}

func main() {
	if len(os.Args[1:]) != 9 {
		stringToRuneAndPrint("Error")
		return
	}
	for _, str := range os.Args[1:] {
		if len(str) != 9 {
			stringToRuneAndPrint("Error")
			return
		}
	}
	matrix := make([][]int, len(os.Args[1:]))
	for i, str := range os.Args[1:] {
		if hasUniqueElements(str) {
			matrix[i] = make([]int, len(str))
			for j, char := range str {
				if char >= '1' && char <= '9' {
					matrix[i][j] = int(char - '0')
				} else {
					matrix[i][j] = 0
				}
			}
		} else {
			stringToRuneAndPrint("Error")
			return
		}
	}
	if hasUniqueSolution(matrix) {
		if solveSudoku(matrix) {
			printGrid(matrix)
		}
	} else {
		stringToRuneAndPrint("Error")
	}
}
