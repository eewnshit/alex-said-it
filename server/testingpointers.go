package main

import "fmt"

func main() {
	string := "yo"
	string_ptr := &string
	*string_ptr = "yo2"

	fmt.Println(string, string_ptr)
}
