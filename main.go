package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i := range args[1:] {
		fmt.Printf("index is %d, value is %s\n", i, args[i])
	}
}
