package Memento

import (
	"log"
	"testing"
)

func Test_Memento(t *testing.T) {
	originator := Originator{}
	careTaker := CareTaker{}
	originator.setState("State #1")
	originator.setState("State #2")
	careTaker.add(originator.saveStateToMemento())
	originator.setState("State #3")
	careTaker.add(originator.saveStateToMemento())
	originator.setState("State #4")

	log.Println("Current State" + originator.state)
	originator.getStateFromMemento(careTaker.get(0))
	log.Println("First Saved State" + originator.state)
	originator.getStateFromMemento(careTaker.get(1))
	log.Println("Second Saved State" + originator.state)
}
