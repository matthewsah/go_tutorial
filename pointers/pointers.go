package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	// attempting to dereference a nil pointer will crash the program
	if message == nil {
		return
	}
	messageStr := *message
	newStr := strings.ReplaceAll(messageStr, "dang", "****")
	*message = newStr
}

func main() {
	// stored at different addresses
	x := 5
	// y := 5

	// z stores the ADDRESS/REFERENCE of x, not the value
	// to get the address/reference we use &
	// z's type is pointer to an integer
	z := &x
	// to access the value at the address, we need to DEREFERENCE it using *
	*z = 6

	fmt.Printf("%d is stored at address %d\n", *z, z)
	fmt.Printf("x = %d\n", x)

	// pointers are often passed as a value into a function so that they can be changed an updated by the function
	// they are also used to improve the performance of the program
	myStr := "dang it"
	myStrReference := &myStr
	removeProfanity(myStrReference)
	fmt.Println(myStr)
}
