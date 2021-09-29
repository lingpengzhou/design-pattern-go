package Builder

import "log"

type Item interface {
	Name() string
	Packing() Packing
	Pricing() float32
}

type Packing interface {
	Pack() string
}

type Wrapper struct {
}

func (wrapper *Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct {
}

func (bottle *Bottle) Pack() string {
	return "Bottle"
}

type Burger struct {
}

func (burger *Burger) Packing() Packing {
	return &Wrapper{}
}

type ColdDrink struct {
}

func (coldDrink *ColdDrink) Packing() Packing {
	return &Bottle{}
}

type VegBurger struct {
	Burger
}

func (vegBurger *VegBurger) Name() string {
	return "vegBurger"
}

func (vegBurger *VegBurger) Pricing() float32 {
	return 25.0
}

type ChickenBurger struct {
	Burger
}

func (chickenBurger *ChickenBurger) Name() string {
	return "ChickenBurger "
}

func (chickenBurger *ChickenBurger) Pricing() float32 {
	return 50.5
}

type Pepsi struct {
	ColdDrink
}

func (pepsi *Pepsi) Name() string {
	return "Pepsi"
}

func (pepsi *Pepsi) Pricing() float32 {
	return 62.12
}

type Coke struct {
	ColdDrink
}

func (coke *Coke) Name() string {
	return "Coke"
}

func (coke *Coke) Pricing() float32 {
	return 31.12
}

type Meal struct {
	items []Item
}

func (meal *Meal) addItem(item Item) {
	//meal.items = make([]Item, 0)
	meal.items = append(meal.items, item)
}

func (meal *Meal) getCost() float32 {
	cost := float32(0.0)
	for _, val := range meal.items {
		cost = cost + val.Pricing()
	}
	return cost
}

func (meal *Meal) showItems() {
	for _, val := range meal.items {
		log.Println("Item:" + val.Name())
		log.Println(",Packing" + val.Packing().Pack())
		log.Println(",Price", val.Pricing())
	}
}

type MealBuilder struct {
}

func (mealBuilder *MealBuilder) prepareVegMeal() Meal {
	var meal Meal
	meal.addItem(&VegBurger{})
	meal.addItem(&Coke{})
	return meal
}

func (mealBuilder *MealBuilder) prepareNonVegMeal() Meal {
	var meal Meal
	meal.addItem(&ChickenBurger{})
	meal.addItem(&Pepsi{})
	return meal
}
