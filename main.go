package main

import (
	"MyProject/simplecalc"
	"fmt"
)

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      int
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

	//Pointers
	i := 8
	j := &i
	fmt.Println("Value of i:", i)
	fmt.Println("Value of j which is address of i:", j)
	fmt.Println("Value pointed by j:", *j)
	*j = 10
	fmt.Println("Updated value of i:", i)

	//Arrays
	nums := [5]int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Println("Index", i, "Value", v)
	}

	ad := Address{
		Address1: " 123 street",
		Address2: " Apt 2",
		State:    "TX",
		Zip:      45627,
	}

	fmt.Println(ad.Address1)
	fmt.Println(ad.Address2)
	fmt.Println(ad.State)
	fmt.Println(ad.Zip)
	fmt.Println(ad)
}
