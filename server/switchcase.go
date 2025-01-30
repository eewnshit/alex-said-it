package main

import (
	"fmt"
)

func print_type_value(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	print_type_value(1)
	print_type_value("teste")
	print_type_value(struct{}{})
}
