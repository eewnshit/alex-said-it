package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(1, 2, 3, add))
	fmt.Println("random number between 1 - 10", rand.Intn(10))
}
