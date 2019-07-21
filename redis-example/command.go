package main

import (
	"github.com/go-redis/redis"
	"fmt"
)


func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.137.18:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	res, err := client.Do("set", "dotest", "testdo").Result()
	if err == redis.Nil {
		fmt.Println("set null")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("ok", res)
	}

	Get := func(redisdb *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		redisdb.Process(cmd)
		return cmd
	}

	v, err := Get(client, "key_does_not_exist").Result()
	fmt.Printf("%q %s\n", v, err)

}

func main() {
	ExampleClient()
}