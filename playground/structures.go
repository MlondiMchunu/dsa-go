package main

import (
	"fmt"
)

type Saiyan struct {
	name  string
	power int
	boss  *Saiyan
}

// composition
type Person struct {
	name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm s\n", p.name)
}

type Saiyan2 struct {
	*Person
	power int
}

func main1() {
	goku := &Saiyan{
		name:  "Goku",
		power: 6000,
	}
	gohan := &Saiyan{
		name:  "Gohan",
		power: 1200,
		boss: &Saiyan{
			name:  "goku",
			power: 1201,
			boss:  nil,
		},
	}

	/*Use composition*/
	goku2 := &Saiyan2{
		Person: &Person{"Goku"},
		power:  9001,
	}
	goku2.Introduce()

	fmt.Println("Name : ", goku.name, "\nPower : ", goku.power)

	Super(goku)
	//goku.power = 6000

	fmt.Println("Name : ", goku.name, "\nPower : ", goku.power)

	fmt.Println("***************\nType *Saiyan receiver of the Super Method\n***************")
	goku.Super()
	fmt.Println(goku.power)

	fmt.Println("***************")
	jeff := NewSaiyan("Jeff", 70000)
	fmt.Println(jeff, "\n", jeff.name)

	fmt.Println("************************")
	fmt.Println(gohan, "\n", gohan.boss)

}

func Super(s *Saiyan) {
	s = &Saiyan{name: "Mlo"}
	s.name = "Mlo"
	s.power = 100
	fmt.Println("Super method ", s.power, " ", s.name)
}

/*
	structures and functions

Here we say the type *Saiyan is the receiver of the Super method
*/
func (s *Saiyan) Super() {
	s.power += 10000
}

/*
*Create a function that returns an instance of the desired type
NOTE: the function doesn't have to return a pointer
*/
func NewSaiyan(Name string, Power int) Saiyan {

	return Saiyan{
		name:  Name,
		power: Power,
	}
}
