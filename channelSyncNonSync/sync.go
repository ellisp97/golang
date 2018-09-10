package main

import (
	"fmt"
	"time"
	"math/rand"
)

// to make goroutine wait it's turn
// message struct that contains channel for reply
type Message struct {
	str string
	wait chan bool
}

// Redeclaration of fan with the new Message struct
func fanIn2(input1, input2 <-chan Message) <-chan Message{
	c := make(chan Message)
	go func(){for {c<- <-input1}}()
	go func(){for {c<- <-input2}}()
	return c
}

// Now with extra channel of bools to set whether new message can be recieved or not
func main(){
	waitForIt := make(chan bool)
	c := fanIn2(boring2(Message{"Joe", waitForIt}), boring2(Message{"Ann", waitForIt}))
	for i:= 0; i< 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

}

func boring2(msg Message) <-chan Message { // returns read only channel of Message structs
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func(){
		for i:= 0; ;i++{
			c <- Message{ fmt.Sprintf("%s %d", msg.str ,i), waitForIt }
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
			<- waitForIt
		}
	}()
	return c
}