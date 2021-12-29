package main

import "fmt"

func main() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
