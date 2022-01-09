package main

import "fmt"

// Generator Pattern is used to generate a sequence of values which is used to produce some output.
// This pattern is widely used to introduce parallelism into loops.
// This allows the consumer of the data produced by the generator to run in parallel
// when the generator function is busy computing the next value.

// Generator func which produces data which might be computationally expensive.
func fib(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	// fib returns the fibonacci numbers lesser than 100
	for i := range fib(100) {
		// Consumer which consumes the data produced by the generator, which further does some extra computations
		v := i * i
		fmt.Println(v)
	}
}
