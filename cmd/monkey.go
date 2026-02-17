package main

import (
	"fmt"
	"os"

	"github.com/cesarleops/zimu/repl"
)

func main() {
	args := os.Args
	fmt.Println("Welcome to monkey ")
	if len(args) == 1 {
		fmt.Println("Starting monkey repl...")
		repl.Start(os.Stdin, os.Stdout)
	}
}
