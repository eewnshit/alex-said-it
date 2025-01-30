package main

import "fmt"

func main() {
	costs := make([]int, 13, 30)
	costs[12] = 12
	fmt.Println(costs)
	fmt.Println(cap(costs))
}
