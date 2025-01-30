package main

import (
	"fmt"
)

const (
	a = "a"
	b = "b"
)

func main() {

	lista := []string{"álex", "fernando", "rocha"}
	fmt.Println(lista[:len(lista)-1])
}
