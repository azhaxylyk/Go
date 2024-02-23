package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
