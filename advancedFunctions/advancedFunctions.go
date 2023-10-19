package main

import "fmt"

// advanced functions
// First class and higher order functions
// A programming language supports first class functions if it allows functions to be passed as data / parameters
// A higher order function is a function that uses a passed in function

// These are often used in HTTP API handlers, Pub/Sub handlers, Onclick callbacks
// pass in the function, don't want to call it when the program runs, but when an event occurs

// keyword defer is unique in go
// usually used for cleanup
// defer <action>
// will perform action right before return

func add(x, y int) (sum int) {
	return x + y
}

func mul(x, y int) (product int) {
	return x * y
}

func aggregate(x, y, z int, arithmetic func(int, int) int) int {
	temp := arithmetic(x, y)
	result := arithmetic(temp, z)
	return result
}

// closures: function that references variables from outside its own function body
// funciton may access and assign to the referenced variables

func concatter() func(string) string {
	doc := ""

	// anonymous functions are lambdas (functions with no name)
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	fmt.Println(aggregate(3, 4, 5, mul))
	fmt.Println(aggregate(1, 2, 3, add))

	fmt.Println("Magic")
	magic := concatter()
	magic("wow")
	magic("this")
	magic("is")
	magic("cool")

	fmt.Println(magic("stuff"))
}
