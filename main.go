package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, a := range args {
		switch a {
		case "-h":
			help()
			return
		}
	}
}

func help() {
	fmt.Printf("mtail - Tail multiples files/directories\n")
	fmt.Printf("\n")
	fmt.Printf("USAGE:\n")
	fmt.Printf("  mtail [flags] <path> [path]...\n")
	fmt.Printf("\n")
	fmt.Printf("FLAGS:\n")
	fmt.Printf("  -h Print this help\n")
}
