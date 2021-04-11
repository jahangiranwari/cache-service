package cache

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

var Expiry time.Duration
var client *redis.Client

func init() {
	exp, err := strconv.Atoi(os.Getenv("CACHE_EXPIRY"))
	if err != nil {
		panic("Missing 'CACHE_EXPIRY' env variable")
	}
	Expiry = time.Duration(exp) * time.Minute
	fmt.Println("Cache expiry set to:", Expiry)

	client = NewClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func NewClient() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}
