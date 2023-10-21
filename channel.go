package main

import "fmt"

func loop(out chan int) {
	for i := 0; i < 2; i++ {
		out <- i
	}
	// close(out)
}

func main2() {
	channel := make(chan int)
	go loop(channel)
	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}
}
