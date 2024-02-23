package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Invalid arguments. Usage: go run . <input_file>.txt <output_file>.txt")
	}
	input, output := os.Args[1], os.Args[2]

	text, err := ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	modifiedContent := ProccessText(text)
	err = WriteFile(output, modifiedContent)
	if err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}
}
