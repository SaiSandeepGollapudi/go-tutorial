package main

import (
	"MyProject/simplecalc"
	"fmt"
)

func main() {
	fmt.Println("Hi world")

	a, b := 6, 4
	fmt.Println(simplecalc.Add(a, b))

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
