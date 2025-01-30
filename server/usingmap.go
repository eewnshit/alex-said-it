package main

import "fmt"

func main() {
	grades := make(map[string]float32, 10)
	grades["álex"] = 6
	grades["letícia"] = 7
	delete(grades, "letícia")

	fmt.Println(grades)
	if _, ok := grades["álex"]; ok {
		fmt.Println("'álex' existe no hashmap")
	}

	chars := []string{"a", "b", "c"}
	for _, char := range chars {
		fmt.Println(char)
	}
}
