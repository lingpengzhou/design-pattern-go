package Memento

type Memento struct {
	state string
}

func SetMementoState(state string) *Memento {
	return &Memento{state: state}
}

func (memento *Memento) getState() string {
	return memento.state
}

type Originator struct {
	state string
}

func (originator *Originator) setState(state string) {
	originator.state = state
}

func (originator *Originator) saveStateToMemento() *Memento {
	return &Memento{
		originator.state,
	}
}

func (originator *Originator) getStateFromMemento(memento Memento) {
	originator.state = memento.getState()
}

type CareTaker struct {
	mementoList []*Memento
}

func (careTaker *CareTaker) add(memento *Memento) {
	careTaker.mementoList = append(careTaker.mementoList, memento)
}

func (careTaker *CareTaker) get(i int) Memento {
	return *careTaker.mementoList[i]
}
