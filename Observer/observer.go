package observertest

import "log"

type Observer interface {
	update(state int)
}

type Subject struct {
	observers []Observer
	state     int
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (subject *Subject) getState() int {
	return subject.state
}

func (subject *Subject) setState(state int) {
	subject.state = state
	subject.notifyAll()
}

func (subject *Subject) attach(observer Observer) {

	subject.observers = append(subject.observers, observer)
}

func (subject *Subject) notifyAll() {
	for _, val := range subject.observers {
		val.update(subject.state)
	}
}

type OriginObserver struct {
	name string
}

func NewOriginObserver() *OriginObserver {
	return &OriginObserver{
		name: "test",
	}
}

func (originObserver *OriginObserver) update(state int) {
	log.Println("originObserver:", state)
}

type PlusObserver struct {
}

func (plusObserver *PlusObserver) update(state int) {
	log.Println("plusObserver:", state+1)
}

type DecreaseObserver struct {
}

func (decreaseObserver *DecreaseObserver) update(state int) {
	log.Println("decreaseObserver:", state-1)
}
