package queue

import (
	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
	"context"
	"log"
	"os"
)

func Publish(path string, channel string) (bool, error){
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	ctx := context.Background()
	rds := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := rds.Publish(
		ctx,
		channel,
		path,
	).Err()

	if err != nil {
		return false, err
	}
	
	return true, nil
}