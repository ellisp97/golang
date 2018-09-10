package main

import (
	"fmt"
	"time"
	"math/rand"
)

//rewrite original fan function using select

				// --- OLD ---
//func fanIn(input1, input2 <-chan string) <-chan string{
//	c := make(chan string)
//	go func(){for {c<- <-input1}}()
//	go func(){for {c<- <-input2}}()
//	return c
//}

				// --- NEW ---  (only ONE goroutine is needed)
func fanInS(input1, input2 <-chan string) <-chan string{
	c := make(chan string)
	go func(){
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
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

func main() {
	c := boring("Joe")
	//to do a total loop timeout
	//timeout := time.After(5*time.Second)
	for{
		select {
		case s := <-c:
			fmt.Println(s)
		case <- time.After(500 * time.Millisecond):
			fmt.Println("Too Slow")
			return
		}
	}
}