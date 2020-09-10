package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 1 argument [http | region].")
		os.Exit(1)
	}

	mode := os.Args[1]

	if mode != "http" && mode != "region" {
		fmt.Println("First argument has to be either 'http' or 'region'.")
		os.Exit(1)
	}

	switch mode {
		case "http":
			startServer()
			break
		case "region":
			startRegion()
			break
	}
}
