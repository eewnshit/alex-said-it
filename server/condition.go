package main

import "fmt"

func main() {
	const age int = 18
	const name string = "Álex Fernando"

	const email string = "eewnshit@gmail.com"

	if age >= 18 {
		fmt.Printf("%s is adult\n", name)
	} else if age < 18 {
		fmt.Printf("HELLL NAHHH, %s is not adult\n", name)
	}

	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	} else {
		fmt.Println("Email is valid")
	}
}
