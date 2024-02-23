package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				PrintStrErr("ERROR: " + err.Error())
				os.Exit(1)
			}
			defer file.Close()

			io.Copy(os.Stdout, file)
		}
	}
}

func PrintStrErr(s string) {
	io.WriteString(os.Stderr, s)
	io.WriteString(os.Stderr, "\n")
}
