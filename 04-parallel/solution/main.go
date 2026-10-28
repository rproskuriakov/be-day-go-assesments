package main

import (
	"fmt"
	"time"
)

func slowSquare(n int) int {
	time.Sleep(time.Second)
	return n * n
}

func main() {
	start := time.Now()
	results := make(chan int)

	for n := 1; n <= 5; n++ {
		go func() {
			results <- slowSquare(n)
		}()
	}

	sum := 0
	for range 5 {
		sum += <-results
	}
	fmt.Println(sum, time.Since(start))
}
