package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)


func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.137.18:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("feekey", "examples", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

	val2, err := client.Get("feekey2").Result()
	if err == redis.Nil {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}

	set, err := client.SetNX("feekey", "value", 10*time.Second).Result()
	fmt.Println(set)

	res, err := client.Do("set", "dotest", "testdo").Result()
	fmt.Println(res)

	res2, err:= client.Append("feekey", "_add").Result()
	fmt.Println(res2)

	val, err = client.Get("feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

}

func main() {
	ExampleClient()
}