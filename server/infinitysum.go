package main

import "fmt"

func sum(nums ...int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		result += num
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(sum(numbers...))
}
