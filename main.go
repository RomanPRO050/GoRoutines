package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go spinner(100 * time.Millisecond)
	n1 := 32
	n2 := 33
	go fmt.Printf("\rFibonacci(%d) = %d\n", n1, fib(n1))
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
