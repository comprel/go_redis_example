package main

import (
	"github.com/go-redis/redis"
	"fmt"
)


func ExampleClient() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.137.18:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisdb.Get("missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}

}

func main() {
	ExampleClient()
}