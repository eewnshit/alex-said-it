package main

import "fmt"

func main() {
	const name string = "Saul Goodman"
	const open_rate float32 = 30.5

	msg := fmt.Sprintf(
		"Hi %s, your open rate is %.1f percent",
		name,
		open_rate,
	)

	fmt.Println(msg)
}
