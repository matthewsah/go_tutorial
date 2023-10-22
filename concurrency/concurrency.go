package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// go doSomething() makes doSomething() execute concurrently with the rest of the code
	// spawns a new goroutine (thread)
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("email received: %s]n", message)
	}()
	fmt.Printf("Email sent: %s\n", message)
}

func main() {
	sendEmail("hello there")
	sendEmail("hello nice to meet you")
	sendEmail("hi")

	// channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
	ch := make(chan int)

	// send data to channel
	// <- is called the channel operator. Data flows in the direction of the arrow and will block until another goroutine is ready to receive the value
	ch <- 69

	// receive data from a channel
	// reads and removes a value from the channel and saves it into variable
	// operation blocks until there is a value in the channel to be read
	num := <-ch

	fmt.Println(num)

	// buffered channels
	// channels can optionally be buffered
	// sending clocks when buffer is full
	// receiving blocks when buffer is empty
	// bufferdCh := make(chan int, 100)

	// channels can be explicitly be closed by a sender
	// close(ch)

	// to check if a channel is closed
	// v, ok := <-ch

	// sending on a closed channelw ill cause the program to panic
	// closing channels is not necessary, though you should close them to indicate explicitly to a receiver that nothing else is going to come across

	// we can iterate through channles using range
	// for item := range ch {
	//     item is next value received from ch
	// }

	// Select statements are used to listen to multiple channels at the same time
	// similar to switch statement but for channels
	// select {
	// case i, ok := <- chInts:
	// 	fmt.Println(i)
	// case s, ok := <- chStrings:
	//     fmt.Println(s)
	// default:
	//     receiving from ch would block
	//     so do something else
	// }

	// Tickers
	// time.Tick() is a standard library function that returns a channel that sends a value on a given interval
	// time.After() sends a value once after the duration has passed
	// time.Sleep() blocks the current goroutine for the specified amount of time

	// a channel can be marked as read-only by casting it from a chan to a <- chan type
	// this can be done by puting param type as <- chan <type>

	// a channel can be marked as write-only by asting it from a chan to a chan <- type
	// this can be done by putting param as chan <- <type>

	// a send to a nil channel blocks forever
	// var c chan string // c is nil
	// c <- "start me!" // blocks

	// a receive from a nil channel blocks forever
	// var c chan string // c is nil
	// fmt.Println(<-c) // blocks

	// A send to a closed channel panics
	// var c - make(chan int, 100)
	// close(c)
	// c <- 1

	// a receive from a closed channel returns zero immediately
	// var c = make(chan int, 100)
	// close(c)
	// fmt.Println(<-c) // 0
}
