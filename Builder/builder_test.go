package Builder

import (
	"log"
	"testing"
)

func Test_Builder(t *testing.T) {
	mealBuilder := MealBuilder{}
	vegMeal := mealBuilder.prepareVegMeal()
	log.Printf("Veg Meal")
	vegMeal.showItems()
	log.Printf("Total Cost:%v", vegMeal.getCost())

	noVegMeal := mealBuilder.prepareNonVegMeal()
	log.Printf("No Veg Meal")
	noVegMeal.showItems()
	log.Printf("Total Cost:%v", noVegMeal.getCost())
}
