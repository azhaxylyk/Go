package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTermSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	parts := strings.Split(strings.TrimSpace(string(out)), " ")

	if len(parts) != 2 {
		log.Fatalln("Unexpected output from stty command")
	}

	cols, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln(err)
	}

	return cols
}

func CheckTerminal(s string) bool {
	termSize := getTermSize()

	if len(s)/8 >= termSize {
		return true
	}

	return false
}
