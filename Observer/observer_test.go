package observertest

import "testing"

func Test_Observer(t *testing.T) {

	originObserver := NewOriginObserver()
	plusObserver := &PlusObserver{}
	decreaseObserver := &DecreaseObserver{}
	subject := NewSubject()
	subject.attach(originObserver)
	subject.attach(plusObserver)
	subject.attach(decreaseObserver)
	subject.setState(1)

}
