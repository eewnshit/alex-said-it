package main

import (
	"errors"
	"fmt"
)

func sub(x int, y int) int {
	return x - y
}

func increment(x int) int {
	x++
	return x
}

func get_point() (x int, y int) {
	return 3, 4
}

func get_coords() (int, int) {
	var x int
	var y int
	return x, y
}

func calculator(a, b int) (mul int, div float32, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = float32(a) / float32(b)
	return mul, div, nil
}

func main() {
	num1 := 10
	num2 := 3
	num1 = increment(num1)

	x, _ := get_point()
	i, j := get_coords()

	mul, div, err := calculator(5, 2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mul)
	fmt.Println(div)

	fmt.Println(x)

	fmt.Println(i, j)

	fmt.Println(num1)
	fmt.Printf("%v\n", sub(num1, num2))
}
