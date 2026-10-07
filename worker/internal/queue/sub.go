package queue

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func Sub(ctx context.Context, channel string) <-chan *redis.Message {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	rds := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	sub := rds.Subscribe(ctx, channel)

	go func() {
		<-ctx.Done()
		sub.Close()
	}()

	return sub.Channel()
}
