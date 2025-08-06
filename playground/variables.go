package main

import (
	"fmt"
)

func mainss() {
	/* var power int
	   power = 9000
	*/
	//*OR* var power int = 9000

	power := getPower()

	fmt.Println("It's over %d\n", power)

	name, power := "mlo", 9001

	fmt.Println("It's also over %d\n", name, power)
}

func getPower() int {
	return 9000
}
