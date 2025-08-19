package main

import (
	"MyProject/simplecalc"
	"fmt"
)

type Address struct {
	Address1 string
	Address2 string
	City   string
	State  string
	Zip    int
}

func main() {
	fmt.Println("Hi world")

	a, b := 6, 4
	fmt.Println(simplecalc.Add(a, b))

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//if else
	num := 7
	if num > 10 {
		fmt.Println("greater")
	} else if num == 10 {
		fmt.Println("equal")
	} else {
		fmt.Println("less")
	}

	//Switch
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday":
		fmt.Println("Tuesday")
	default:
		fmt.Println("Weekend")
	}

	//defer
	defer fmt.Println("Hi")
	fmt.Println("world")
	defer fmt.Println("go")

	a := Address{
		Address1: " 123 street",
		Address2: " Apt 2",
		State:    "TX",
		Pin:      45627,
	}
	fmt.Println(a.Address1)
	fmt.Println(a.Address2)
	fmt.Println(a.State)
	fmt.Println(a.Pin)
	fmt.Println(a)
	}

}
