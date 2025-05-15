package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

func greetPerson(person Person) {
	fmt.Println("My name is", person.name, "My age is", person.age, "My gender is", person.gender)
}
func main() {
	me := Person{
		name:   "Sankalp Pol",
		age:    42,
		gender: "Male",
	}
	fmt.Println("My name is", me.name, "My age is", me.age, "My gender is", me.gender)

	me.name = "Kunal"

	fmt.Println("My name is", me.name, "My age is", me.age, "My gender is", me.gender)

	you := Person{
		name:   "Rahul Jaykar",
		age:    64,
		gender: "Male",
	}
	fmt.Println(you)

	greetPerson(me)
	greetPerson(Person{
		name:   "Kaveri",
		age:    77,
		gender: "Female",
	})
}
