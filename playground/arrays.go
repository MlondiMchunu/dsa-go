package main

import "fmt"

func main2() {

	var scores [10]int
	scores[0] = 5
	scores[3] = 7

	marks := [5]int{4, 3, 24, 3, 67}

	for i := 0; i < len(marks); i++ {
		fmt.Print(marks[i], "\t")
	}
	fmt.Println()
	for i := 0; i < len(scores); i++ {
		fmt.Print(scores[i], " ")
	}
}
