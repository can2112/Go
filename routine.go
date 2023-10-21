package main

import "fmt"

// "time"

func print(a string, out chan bool) {
	for i := 0; i < 3; i++ {
		fmt.Println(a, "*", i)
	}
	out <- true
}
func main3() {
	channel := make(chan bool)
	go print("go-routine", channel)
	<-channel
}


