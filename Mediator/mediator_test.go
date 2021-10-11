package Mediator

import "testing"

func Test_Mediator(t *testing.T) {
	robert := User{"Robert"}
	john := User{"John"}
	robert.sendMessage("Hi! Robert!")
	john.sendMessage("Hi! John")
}
