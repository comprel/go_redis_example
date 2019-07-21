package redigo_example

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"flag"
)


var (
	pool *redis.Pool
	//redisServer = flag.String("192.168.137.18", ":6379", "")
)


func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}


func Get() redis.Conn {
	return pool.Get()
}

func main() {
	flag.Parse()
	pool = newPool("192.168.137.18:6379")
	connections := pool.Get()
	defer connections.Close()

	set_res, err := connections.Do("SET", "new_test", "redigo")
	if err != nil {
		fmt.Println("err while set key :", err)
	}else {
		fmt.Println(set_res)
	}

	is_exists, err := redis.Bool(connections.Do("EXISTS", "new_test"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(is_exists)
	}
	get_res, err := redis.String(connections.Do("GET", "new_test"))
	if err != nil {
		fmt.Println("get err:", err)
	} else {
		fmt.Println(get_res)
	}

}