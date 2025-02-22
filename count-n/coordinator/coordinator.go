package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
	"time"
)


type Task struct {
	ID uint64
	Start uint64
	End uint64
}

type Result struct {
	TaskID uint64
	Sum uint64
}


func main() {
	numWorkers := 10

	totalNumbers := uint64(100000000)

	chunkSize := totalNumbers / uint64(numWorkers)

	results := make([]Result, numWorkers)

	var wg sync.WaitGroup

	startTime := time.Now()

	for i:=0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID uint64) {
			defer wg.Done()

			client, err := rpc.Dial("tcp", "localhost:8081")
			if err != nil {
				log.Fatal("Failed to dial worker:", err)
			}

			task := Task{
				ID: workerID,
				Start: workerID * chunkSize + 1,
				End: (workerID + 1) * chunkSize,
			}

			var result Result

			err = client.Call("WorkerService.AssignTask", task, &result)
			if err != nil {
				log.Fatal("Failed to call WorkerService.AssignTask:", err)
			}

			results[workerID] = result
		}(uint64(i))
	}

	wg.Wait()

	totalSum := uint64(0)
	for _, result := range results {
		totalSum += result.Sum
	}

	endTime := time.Now()

	fmt.Printf("Total sum: %d\n", totalSum)
	fmt.Printf("Time taken: %s\n", endTime.Sub(startTime))
}