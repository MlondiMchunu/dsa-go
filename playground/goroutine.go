package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
}

func printNumbers() {

	for i := 1; i <= 26; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Number : ", i)
	}
}

func printLetters() {
	for i := 'A'; i <= 'Z'; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Letter : ", i)
	}
}
