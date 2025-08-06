package main

import (
	"fmt"
	"os"
)

func mai() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
