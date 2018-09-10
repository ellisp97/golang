package main

import "fmt"

// Here GoRotuines/Concurrency/Parellelism is useful when high processing is required
// If I need to process thousands of factorial calculations, Put them in a goroutine
// ,running them concurrently and in parallel will help speed up process by maximising CPU cores

const input = 10

func main() {
	in := gen(input)
	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

// Currently set to do 0-10 factorial numbers from constant
func gen(input int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < input; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out:= make( chan int)
	go func() {
		for n := range in{
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int{
	total := 1
	for i:=n; i>0; i--{
		total *= i
	}
	return total
}