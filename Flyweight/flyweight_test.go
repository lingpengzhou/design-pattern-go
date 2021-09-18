package Flyweight

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_FlyWeight(t *testing.T) {
	colors := [...]string{"Red", "Green", "Blue", "White", "Black"}
	shapeFactory := ShapeFactory{
		circleMap: make(map[string]Shape, 0),
	}
	for i := 0; i < 20; i++ {
		v := getRand(4)
		circle := shapeFactory.getCircle(colors[v])
		circle.SetXYRadius(getRand(500), getRand(600), getRand(900))
		circle.draw()
	}
}

func getRand(num int) int {
	rand.Seed(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}
