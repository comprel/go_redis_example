package redigo_example

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func echoReceive(res interface{}, err error){
	if err != nil{
		fmt.Println(err)
	}else {
		if res != nil{
			fmt.Printf("---------  ")
			fmt.Println(string(res.([]byte)))
		}else {
			fmt.Println(res)
		}

	}
}

func main() {
	c1, err := redis.Dial("tcp", "192.168.137.18:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	//
	//c1.Send("Get", "my_test")
	//c1.Flush()
	//echoReceive(c1.Receive())
	//
	//c1.Send("Get", "my_test2")
	//c1.Flush()
	//echoReceive(c1.Receive())

	var value1 string
	var value2 string
	var value3 string

	c1.Send("MULTI")
	c1.Send("Get", "my_test")
	c1.Send("Get", "my_test2")
	c1.Send("Get", "my_test3")
	r, err := redis.Values(c1.Do("EXEC"))
	//fmt.Println(r)
	if _, err := redis.Scan(r, &value1, &value2, &value3); err == nil {
		fmt.Println(value1)
		fmt.Println(value2)
		fmt.Println(value3)
	}
}