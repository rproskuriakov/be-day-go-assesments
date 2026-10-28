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
	sum := 0
	for n := 1; n <= 5; n++ {
		sum += slowSquare(n)
	}
	fmt.Println(sum, time.Since(start)) // 55 5s
}
