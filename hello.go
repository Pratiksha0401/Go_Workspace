package main

import (
	"fmt"
	"os"
)

func main() {
	//first arg is always program name
	fmt.Println("len", len(os.Args))
	//so start from index 1
	fmt.Println("len", len(os.Args[1:]))

	for _, arg := range os.Args[1:] {
		fmt.Println("Hello ",arg,"!")
	}

	fmt.Println("\nArguments:")
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("\tArgument[%d]: %s\n", i, os.Args[i])
	}
}