package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("OS args:")
	arguments := os.Args
	fmt.Printf("Full array: %v\n", arguments)
	for index, arg := range arguments {
		fmt.Printf("idx: %d element: %s\n", index, arg)
	}
}
