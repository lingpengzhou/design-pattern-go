package Flyweight

import "log"

type Shape interface {
	draw()
	SetXYRadius(x int, y int, radius int)
}

type Circle struct {
	x, y, radius int
	color        string
}

func (circle *Circle) SetXYRadius(x int, y int, radius int) {
	circle.x = x
	circle.y = y
	circle.radius = radius
}

func (circle *Circle) draw() {
	log.Printf("Circle is Drawing : Color->%v,X->%v,y->%v", circle.color, circle.x, circle.y)
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

func (shapeFactory *ShapeFactory) getCircle(color string) Shape {
	circle := shapeFactory.circleMap[color]
	if circle == nil {
		circle = &Circle{color: color}
		shapeFactory.circleMap[color] = circle
		log.Println("Creating circle of color : " + color)
	}
	return circle
}
