package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}

}

//===========ORDER=OF=OPS==================================

//-----------in-go-routine-------------------------------
// recieve chan of ints
// launch go routine code flows straight through - out of the way and running
// in chan of ints
// 100 times going to assign 3-13 onto out
// when it runs its gonna assign 3 and its going to HOLT - put a value on a channle and the code will STOP

//-----------factorial-go-routine-----------------------
// recieves in
// launch another go routine which is going to be off and running
// return out which is a channel of int and assign this to c0
// flows all the way to c9

// have some var y ints
// loop over function merge with vars c0-c9 which are recieving chan of ints with a value coming off of the channel from factorial()

// in factorial now we're gonna range over in (chan of ints) - THIS PULLS A VALUE OFF OF THE IN CHANNEL (the one which got passed in) e.g. factorial(in)
// This came from gen() which puts a value onto the channel, SO WHEN VALUE GETS PULLED OFF - (out <- j) - ANOTHER CAN BE STUCK ON
// Thus goes onto next range in (same channel) out<-fact(n) can pull another value off of that channel passing that onto out which is the specific c0-c9 value


//--------------merge--------------
// cs has unlimited number of type chan ints

// function output where they get processed , recieve a chan of ints , range over it, take value off and put it onto out channel
// out channel is a chan of ints, we take many channels assign to cs, many channels values will be pulled off
// output function does this , range over all of our channels, launch goroutine putting channel in there
// ranging over that channel get value and put it onto out

// close out thus range knows channels closed we're done with the program

//===============================================================================

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}


func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
