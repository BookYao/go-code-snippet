package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(num int) int {
	if num < 2 {
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s Num", os.Args[0])
		os.Exit(0)
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: param error!")
		os.Exit(0)
	}

	go spinner(100 * time.Millisecond)
	fibResult := fibonacci(num)
	fmt.Println("fibResult:", fibResult)
}
