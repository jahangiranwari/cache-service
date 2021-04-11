package task

import (
	"fmt"
	"log"
	"os"

	"github.com/hibiken/asynq"
)

var client *asynq.Client

func init() {
	fmt.Println("Initializing Asynq...")
	redisURL := os.Getenv("REDIS_URL")

	rconn, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	options, _ := rconn.(asynq.RedisClientOpt)
	client = asynq.NewClient(options)
}

func GetClient() *asynq.Client {
	return client
}
