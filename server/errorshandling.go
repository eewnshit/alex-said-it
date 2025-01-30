package main

import (
	"errors"
	"fmt"
)

type User struct {
	name         string
	phone_number int
}

func get_user() (User, error) {
	user := User{
		name:         "álex",
		phone_number: 71986758123,
	}
	return user, errors.New("teste")
}

func main() {
	alex, err := get_user()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(alex)
}
