package jobs

import (
	"go-async-training-worker/internal/db"
	"log"
	"fmt"
)

func Write(c chan [][]string) {
	fmt.Println("writing")
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
