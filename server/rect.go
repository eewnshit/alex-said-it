package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{
		width:  10,
		height: 5,
	}
	fmt.Println(r.area()) // output: 50
}
