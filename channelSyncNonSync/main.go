package main

import (
	"fmt"
	"time"
	"math/rand"
)


// Allows independent calls from Joe or Ann for non sequential output

func main(){
	c := fanIn(boring("Joe"), boring("Ann"))
	for i :=0 ; i <10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Both boring goodbye")
}

// Without fanIn Joe and Ann would take turns calling the boring func
func fanIn(input1, input2 <-chan string) <-chan string{
	c := make(chan string)
	go func(){for {c<- <-input1}}()
	go func(){for {c<- <-input2}}()
	return c
}

func boring(msg string) <-chan string { // returns read only channel of strings
	c := make(chan string)
	go func(){
		for i:= 0; ;i++{
			c <- fmt.Sprintf("%s %d", msg ,i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}