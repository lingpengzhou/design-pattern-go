package Mediator

import (
	"log"
	"time"
)

type ChatRoom struct {
}

func (chatRoom *ChatRoom) showMessage(user *User, message string) {
	log.Println(time.Now())
	log.Println(user.Name)
	log.Println(message)
}

type User struct {
	Name string
}

func setUser(name string) *User {
	return &User{
		name,
	}
}

func (user *User) sendMessage(message string) {
	chatRoom := ChatRoom{}
	chatRoom.showMessage(user, message)
}
