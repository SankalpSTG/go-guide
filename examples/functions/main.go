package main

import "fmt"

func greetings(greeting string, name string, age int) {
	fmt.Println(greeting, name, "your age is", age)
}
func greetings2(greeting, name string, age int) {
	fmt.Println(greeting, name, "your age is", age)
}
func add(a, b int) int {
	return a + b
}
func moodyAdder(a, b int) (string, int) {
	if a%2 == 0 && b%2 == 0 {
		return "Happy", a + b
	} else {
		return "Sad", a - b
	}
}
func moodyAdder2(a, b int) (mood string, total int) {
	if a%2 == 0 && b%2 == 0 {
		mood = "Happy"
		total = a + b
	} else {
		mood = "Sad"
		total = a - b
	}
	return
	// OR
	// return mood, total
}
func factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func main() {
	greetings("Welcome", "Rahul", 34)

	greetings2("Welcome", "Rahul", 34)

	result := add(4, 5)
	fmt.Println(result)

	mood, result := moodyAdder(4, 5)
	fmt.Println(mood, result)

	mood, _ = moodyAdder(10, 4)
	fmt.Println(mood)

	mood, value := moodyAdder2(10, 4)
	fmt.Println(mood, value)

	fact := factorial(7)
	fmt.Println(fact)
}
