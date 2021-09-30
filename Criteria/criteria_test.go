package Criteria

import "testing"

func Test_Criteria(t *testing.T) {
	persons := make([]Person, 0)
	persons = append(persons, *setPerson("Robert",
		"Male",
		"Single"))
	persons = append(persons, *setPerson("John", "Male", "Married"))
	persons = append(persons, *setPerson("Diana", "Female", "Single"))
	persons = append(persons, *setPerson("Laura", "Female", "Married"))
	persons = append(persons, *setPerson("Mike", "Male", "Single"))
	persons = append(persons, *setPerson("Bobby", "Male", "Single"))
	male := CriteriaMale{}
	female := CriteriaFemale{}
	single := CriteriaSingle{}
	singleMale := SetAndCriteria(&single, &male)
	singleOrFemale := OrCriteria{&single, &male}
	println("Males:")
	printPersons(male.meetCriteria(persons))
	println("Females:")
	printPersons(female.meetCriteria(persons))
	println("Single Males:")
	printPersons(singleMale.meetCriteria(persons))
	println("Single Or Females:")
	printPersons(singleOrFemale.meetCriteria(persons))
}
