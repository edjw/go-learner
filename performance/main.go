package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // Get the current time

	var n int = 0
	for i := 0; i < 100000000; i++ {
		n += 1
	}

	elapsed := time.Since(start)
	fmt.Println("Function took :", elapsed)

}

// this is much faster (~2609ms) than the other python one (~5400ms). no idea why. but still much slower than the go one (61ms)
