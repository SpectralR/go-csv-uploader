package main

import (
	"context"
	"sync"

	"go-async-training-worker/internal/jobs"
	"go-async-training-worker/internal/queue"
)

func main() {
	file := <-queue.Sub(context.Background(), "fileParse")

	var wg sync.WaitGroup
	parsed := make(chan [][]string)
	valid := make(chan [][]string)

	wg.Add(3)

	go func(){ 
		defer close(parsed)
		defer wg.Done()
		jobs.Parse(file.Payload, parsed) 
	}()
	go func(){
		defer close(valid)
		defer wg.Done()
		jobs.Validate(parsed, valid)
	}()
	go jobs.Write(valid, &wg)

	wg.Wait()
}
