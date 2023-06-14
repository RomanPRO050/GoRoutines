package main

import (
	"fmt"
	"time"
)

func main() {

	go spinner(100 * time.Millisecond)
	n1 := 10
	n2 := 15
	go fmt.Printf("\rFibonacci(%d) = %d\n", n1, fib(n1*2))
	go fmt.Printf("\rFibonacci(%d) = %d\n", n2, fib(n2))
	time.Sleep(time.Second)
	//fmt.Printf("\rFibonacci(%d) = %d\n")

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}
