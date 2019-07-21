package main

import (
	"github.com/go-redis/redis"
)


func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.137.18:6379",
	})
	//rdb.AddHook()

	rdb.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Ping()
		pipe.Ping()
		return nil
	})
}

func main() {
	ExampleClient()
}