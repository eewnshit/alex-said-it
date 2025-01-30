package main

import (
	"fmt"
)

type message_send struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

type admin struct {
	user
	permissions []string
}

func can_send_message(m_to_send message_send) bool {
	if m_to_send.sender.name == "" || m_to_send.recipient.name == "" {
		return false
	}

	if m_to_send.sender.number == 0 || m_to_send.recipient.number == 0 {
		return false
	}
	return true
}

func main() {

	alex := user{
		name:   "álex",
		number: 71953428166,
	}

	admin := admin{
		permissions: []string{"read", "write", "delete"},
		user:        alex,
	}

	fmt.Println(admin.name, admin.permissions)

	carla := user{
		name:   "carla",
		number: 71952390578,
	}

	new_message := message_send{
		message:   "eae parceira",
		sender:    alex,
		recipient: carla,
	}

	fmt.Println(can_send_message(new_message))
}
