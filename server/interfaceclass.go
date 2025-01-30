package main

import (
	"fmt"
	"time"
)

func send_message(msg message) {
	fmt.Println(msg.get_message())
	fmt.Println("==================================")
}

type message interface {
	get_message() string
}

type birthday_message struct {
	birthday_time  time.Time
	recipient_name string
}

func (bm birthday_message) get_message() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipient_name, bm.birthday_time.Format(time.RFC3339))
}

func (sr sending_report) get_message() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.report_name, sr.number_of_sends)
}

type sending_report struct {
	report_name     string
	number_of_sends int
}

func test(m message) {
	send_message(m)
	fmt.Println("====================================")
}

func main() {
	test(sending_report{
		report_name:     "First Report",
		number_of_sends: 10,
	})
	test(birthday_message{
		recipient_name: "John Doe",
		birthday_time:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sending_report{
		report_name:     "First Report",
		number_of_sends: 10,
	})
	test(birthday_message{
		recipient_name: "Bill Deer",
		birthday_time:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}
