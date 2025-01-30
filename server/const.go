package main

import (
	"fmt"
	"strings"
)

func main() {

	const first_name = "alex"
	const last_name = "fernando"
	const full_name = first_name + " " + last_name
	fmt.Printf("%q\n", strings.Title(full_name))
}
