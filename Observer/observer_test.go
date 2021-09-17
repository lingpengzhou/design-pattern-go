package observertest

import "testing"

func Test_Observer(t *testing.T) {

	originObserver := NewOriginObserver()
	plusObserver := NewPlusObserver()
	decreaseObserver := NewDecreaseObserver()
	subject := NewSubject()
	subject.attach(originObserver)
	subject.attach(plusObserver)
	subject.attach(decreaseObserver)
	subject.setState(1)

}
