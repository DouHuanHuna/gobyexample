package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	rdb.Set(context.Background(), "key1", "value1", 0)
	rdb.Set(context.Background(), "key2", "value2", 0)
	rdb.Set(context.Background(), "key3", "value3", 0)

	keys := []string{"key1", "key2", "key3"}
	values, err := rdb.MGet(context.Background(), keys...).Result()
	if err != nil {
		log.Fatalf("Error fetching values: %v", err)
	}

	fmt.Println(values)
}
