package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	// example of a golang while loop
	for i := 0; ; i++{
		// send
		c <- "ping"
	}
}

func print(c chan string) {
	for {
		// recieve
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// need to use make to initialise any channel type
	var c chan string = make(chan string)

	go ping(c)
	go print(c)

	fmt.Scanln(<-c)
}