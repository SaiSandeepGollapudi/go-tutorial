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

type Triangle struct {
	width  float64
	height float64
}

func (t Triangle) area() float64 {
	return (t.width * t.height) / 2
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

	//Structs
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

	//slice
	s := make([]int, 2, 3) // 2= len, 3 capacity
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	s = append(s, 6)
	s = append(s, 7)

	fmt.Println("s after appending", s)

	//slice with same len and capacity
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	slice := q[1:3]

	fmt.Println(slice)

	//Map

	food := map[string]string{
		"Avacado":   "Guac",
		"Blueberry": "Pie",
	}

	food["carrot"] = "cake"
	fmt.Println("after adding new food:", food)

	if val, ok := food["Avacado"]; ok {
		fmt.Println("Avacado  is present ", val)

	} else {
		fmt.Println("Avacado is not present")

	}

	for key, value := range food {
		fmt.Println(key, value)
	}

	//Method call of Triangle
	tri := Triangle{2, 4}
	fmt.Println("Area:", tri.area())

}
