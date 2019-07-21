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


	//  尝试自定义boolcmd

	GetBool := func(redisdb *redis.Client, key string) *redis.BoolCmd {
		cmd := redis.NewBoolCmd("get", key)
		redisdb.Process(cmd)
		return cmd
	}
	SetBool := func(redisdb *redis.Client, key string, value bool) *redis.BoolCmd {
		cmd := redis.NewBoolCmd("set", key, value)
		redisdb.Process(cmd)
		return cmd
	}
	SetBool(client, "dobool", true)
	v2, err := GetBool(client, "dobool").Result()
	fmt.Println(v2, err)

	// 返回命令列表
	res3 := client.Command().String()
	fmt.Println(res3)

	GetDuration := func(redisdb *redis.Client, value time.Duration) *redis.DurationCmd {
		cmd := redis.NewDurationCmd(value)
		redisdb.Process(cmd)
		return cmd
	}

	v2Duration, err := GetDuration(client, 500 * time.Millisecond).Result()
	fmt.Println(v2Duration, err)
}

func main() {
	ExampleClient()
}