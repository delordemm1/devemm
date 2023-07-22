package main

import (
	"fmt"
	"os"

	"github.com/delordemm1/devemm-go/internal/cmd"
)

func main() {

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "web":
		if len(os.Args) < 3 {
			printUsage()
			os.Exit(1)
		}
		switch os.Args[2] {
		case "serve":
			cmd.WebServe()
		default:
			printUsage()
			os.Exit(1)
		}

	default:
		printUsage()
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  web serve - Run web server")
}
