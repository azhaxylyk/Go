package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 || args[0] != "-c" {
		fmt.Printf("Usage: %s -c <num> <file1> [<file2>...]\n", os.Args[0])
		os.Exit(1)
	}

	num := args[1]

	for _, filename := range args[2:] {
		file, err := os.Open(filename)
		if err != nil {
			PrintError(filename)
			continue
		}
		defer file.Close()

		fileInfo, _ := file.Stat()
		fileSize := fileInfo.Size()

		offset := fileSize - int64(len(num))
		if offset < 0 {
			offset = 0
		}

		buffer := make([]byte, offset)
		_, _ = file.ReadAt(buffer, offset)

		PrintFileHeader(filename)
		fmt.Printf("%s\n\n", buffer)
	}
}

func PrintError(filename string) {
	fmt.Printf("==> %s <==\n", filename)
	fmt.Printf("open %s: no such file or directory\n\n", filename)
}

func PrintFileHeader(filename string) {
	fmt.Printf("==> %s <==\n", filename)
}
