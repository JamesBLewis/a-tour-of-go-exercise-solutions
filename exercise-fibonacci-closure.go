package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	counter := 0
	fib1 := 0
	fib2 := 1
	return func() int {
		if counter < 2 {
			counter++
			return counter - 1
		}
		result := fib1 + fib2
		fib1 = fib2
		fib2 = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
