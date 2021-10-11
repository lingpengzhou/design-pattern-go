package Criteria

import (
	"log"
	"reflect"
)

type Criteria interface {
	meetCriteria(persons []Person) []Person
}

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func setPerson(name string, gender string, maritalStatus string) *Person {
	return &Person{
		name, gender, maritalStatus,
	}
}

type CriteriaMale struct {
}

func (criteriaMale *CriteriaMale) meetCriteria(persons []Person) []Person {
	malePersons := make([]Person, 0)
	for _, val := range persons {
		if val.gender == "Male" {
			malePersons = append(malePersons, val)
		}
	}
	return malePersons
}

type CriteriaFemale struct {
}

func (criteriaFemale *CriteriaFemale) meetCriteria(persons []Person) []Person {
	femalePersons := make([]Person, 0)
	for _, val := range persons {
		if val.gender == "Female" {
			femalePersons = append(femalePersons, val)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {
}

func (criteriaSingle *CriteriaSingle) meetCriteria(persons []Person) []Person {
	singlePersons := make([]Person, 0)
	for _, val := range persons {
		if val.maritalStatus == "Single" {
			singlePersons = append(singlePersons, val)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func SetAndCriteria(criteria Criteria, otherCriteria Criteria) *AndCriteria {
	return &AndCriteria{
		criteria, otherCriteria,
	}
}

func (andCriteria *AndCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaPersons := andCriteria.criteria.meetCriteria(persons)
	return andCriteria.otherCriteria.meetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func SetOrCriteria(criteria Criteria, otherCriteria Criteria) *OrCriteria {
	return &OrCriteria{
		criteria, otherCriteria,
	}
}

func (orCriteria *OrCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaPersons := orCriteria.criteria.meetCriteria(persons)
	otherCriteriaPersons := orCriteria.otherCriteria.meetCriteria(firstCriteriaPersons)

	for _, val := range otherCriteriaPersons {
		if !contains(firstCriteriaPersons, val) {
			firstCriteriaPersons = append(firstCriteriaPersons, val)
		}
	}
	return firstCriteriaPersons
}

func contains(person []Person, s Person) bool {
	for _, p := range person {
		if reflect.DeepEqual(p, s) {
			return true
		}
	}
	return false
}

func printPersons(persons []Person) {
	for _, val := range persons {
		log.Print("Person :[Name:" + val.name +
			", Gender : " + val.gender + ", Marital Status :" + val.maritalStatus)
	}
}
