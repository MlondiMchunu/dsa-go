package main

import "fmt"

func mainz() {
	slice := []int{3, 5, 7, 9, 1, 4, 2}
	lookup := make(map[string][]int)
	lookup["goku"] = slice
	power, exist := lookup["gok"]

	total := len(slice)

	fmt.Println(power, exist)
	fmt.Printf("%v", total)

	//when you need a map as a field of a structure, define it as:
	type Saiyan struct {
		Name    string
		Friends map[string]*Saiyan
	}

	//initialize the above
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	//goku.Friends["krillin"]=...
}
