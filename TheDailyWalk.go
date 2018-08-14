package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

// Challenge : Every *morning* Alice and Bobo go for a walk following the same routine every day
// 			 : first, they prepare grabbing sunglasses, a belt , close/open windows, turn off fans, pocketing keys
//			 : once ready, takes about 60-90 seconds, they arm the alarm, which has 60 second delay
//			 : while alarm they put on shoes, takes about 35-45 sec
// 			 : leave house together before alarm has finished
//	Similuate this with a Program

func doThing(name string, thing string, timeTaken int, timeOffset int, c chan int){
	fmt.Println(name + " started " + thing)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dur := timeTaken + r.Intn(timeOffset)
	time.Sleep(time.Duration(dur*10) * time.Millisecond)
	fmt.Println(name + " spent " + strconv.Itoa(dur) + " seconds " + thing)
	c <- dur
}

func getReady(name string, c chan int) {
	doThing(name, "getting ready", 60, 30, c)
}

func putShoes(name string, c chan int){
	doThing(name, "putting on shoes", 35, 10, c)
}

func armAlarm(c chan int){
	time.Sleep(300 * time.Millisecond)
	fmt.Println("tick .. tick .. tick ..")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Alarm is armed.")
	c <- 1
}

func main () {
	fmt.Println("Let's go for a walk!")
	c := make(chan int)
	go getReady("Bob", c)
	go getReady("Alice", c)
	_, _ = <-c, <-c
	alarmChan := make(chan int)
	fmt.Println("Alarming alarm.")
	go armAlarm(alarmChan)
	go putShoes("Bob", c)
	go putShoes("Alice", c)
	_, _ = <-c, <-c
	fmt.Println("Exiting and Locking the door")
	_ = <-alarmChan
}