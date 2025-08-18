package main

import (
	"MyProject/simplecalc"
	"fmt"
)

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

	//defer
	defer fmt.Println("Hi")
	fmt.Println("world")
	defer fmt.Println("go")

}
