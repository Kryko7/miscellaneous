package main

import (
	"fmt"
	"time"
)


func main() {

	totalNumbers := uint64(10000000000)

	startTime := time.Now()
	
	totalSum := uint64(0)

	for i:= 1; i <= int(totalNumbers); i++ {
		totalSum += uint64(i)
	}

	endTime := time.Now()

	fmt.Printf("Total sum: %d\n", totalSum)
	fmt.Printf("Time taken: %s\n", endTime.Sub(startTime))
}