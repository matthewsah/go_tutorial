package main

import "fmt"

type messageToSend struct {
	message    string
	sender     string
	receipient string
}

type car struct {
	make  string
	model string
	wheel struct {
		radius   int
		material string
	}
}

type truck struct {
	car
	bedSize int
}

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	x := messageToSend{}
	x.message = "hello"
	x.sender = "matt"
	x.receipient = "alyssia"

	myCar := car{}
	myCar.make = "toyota"
	myCar.model = "prius"
	myCar.wheel.radius = 3
	myCar.wheel.material = "rubber"

	myTruck := truck{}
	myTruck.make = "toyota"
	myTruck.model = "prius"
	myTruck.wheel.radius = 3
	myTruck.wheel.material = "rubber"
	myTruck.bedSize = 15

	r := rect{}
	r.height = 5
	r.width = 4

	fmt.Println(myCar)
	fmt.Println(myTruck)
	fmt.Println(r.area())
}
