package redigo_example

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	setdb := redis.DialDatabase(12)
	setPasswd := redis.DialPassword("")


	c1, err := redis.Dial("tcp", "192.168.137.18:6379", setdb, setPasswd)
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	c2, err := redis.Dial("tcp", "192.168.137.18:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c2.Close()

	resset, err := redis.String(c1.Do("SET", "my_test", "redigo"))
	if err != nil {
		fmt.Println("set err")
	} else {
		fmt.Println(resset)
	}

	//获取value并转成字符串
	account_balance, err := redis.String(c1.Do("GET", "my_test"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(account_balance)
	}

	c2_get, err := redis.String(c2.Do("GET", "my_test"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(c2_get)
	}
}