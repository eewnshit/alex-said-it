package main

import "fmt"

func main() {
	// var smsSendingLimit int = 50
	// var costPerSMS float64 = 0.042
	// var hasPermission bool = true
	// var username string = "álex"

	smsSendingLimit := 50
	hasPermission := true
	costPerSMS := 0.042
	username := "álex"

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

}
