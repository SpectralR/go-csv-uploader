package main

import (
	"context"

	"go-async-training-worker/internal/jobs"
	"go-async-training-worker/internal/queue"
)

func main() {
	file := <-queue.Sub(context.Background(), "fileParse")

	channel := make(chan [][]string)

	go jobs.Parse(file.Payload, channel)
	go jobs.Validate(channel)
	go jobs.Write(channel)
}
