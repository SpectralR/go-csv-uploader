package jobs

import (
	"go-async-training-worker/internal/db"
	"log"
	"sync"
)

func Write(c chan [][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	channel := <-c
	if channel != nil {
		db.Init(".env")
		query := db.PrepareInsert(channel)
		_, err := db.DB.Query(query)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}
