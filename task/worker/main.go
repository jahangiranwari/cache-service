package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/jahangiranwari/cache-service/task"
)

func main() {
	fmt.Println("Initializing Asynq worker...")

	redisURL := os.Getenv("REDIS_URL")
	rconn, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	options, _ := rconn.(asynq.RedisClientOpt)
	server := asynq.NewServer(options, asynq.Config{
		Concurrency: 4,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.UpdateCache, task.HandleUpdateCacheTask)
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
