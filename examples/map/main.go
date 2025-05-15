package main

import "fmt"

func main() {
	var person = map[string]string{"name": "Sankalp Pol", "gender": "Male"}
	fmt.Println(person)

	fmt.Println(person["name"])

	person["name"] = "Kunal"  // this updates name
	person["city"] = "Mumbai" // this adds a key "city" with value "Mumbai"

	fmt.Println(person)

	delete(person, "name") //name will be deleted

	fmt.Println(person)

	name, ok := person["name"]
	if !ok {
		fmt.Println("Name doesn't exist")
	} else {
		fmt.Println("Name exists as", name)
	}

	city, ok := person["city"]
	if !ok {
		fmt.Println("City doesn't exist")
	} else {
		fmt.Println("City exists as", city)
	}

	for key, value := range person {
		fmt.Println(key, value)
	}

}
