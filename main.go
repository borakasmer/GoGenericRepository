package main

import (
	"fmt"
	"strconv"
)

func main() {
	var personRepository = Repository[*Person]{}
	var person = new(Person)
	person.Name = "Bill"
	person.Surname = "Gates"

	personRepository.AddEntity(person)
	personRepository.AddEntity(&Person{Name: "Bora", Surname: "Kasmer"})
	personRepository.AddEntity(&Person{Name: "Duru", Surname: "Kasmer"})

	firstPerson := personRepository.GetOne()
	fmt.Println("Get One:" + firstPerson.GetFullName())

	fmt.Println("Person List")
	fmt.Println("-----------------------")
	for _, per := range personRepository {
		fmt.Println(per.GetFullName())
	}

}

type Repository[T IBaseEntity] []T

func (entitySet *Repository[IBaseEntity]) AddEntity(entity IBaseEntity) {
	*entitySet = append(*entitySet, entity)
}

func (entitySet *Repository[IBaseEntity]) GetOne() IBaseEntity {
	first := (*entitySet)[0]
	*entitySet = (*entitySet)[1:]
	return first
}

type IBaseEntity interface {
	GetFullName() string
}
type Player struct {
	No   int
	Nick string
}

func (pl *Player) GetFullName() string {
	return pl.Nick + " : " + strconv.Itoa(pl.No)
}

type Person struct {
	Name    string
	Surname string
}

func (ps *Person) GetFullName() string {
	return ps.Name + " : " + ps.Surname
}
