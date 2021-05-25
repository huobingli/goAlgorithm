package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "47.114.171.118:6380")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("set", "key1", "value1")
	// 若操作失败则返回
	if err != nil {
		fmt.Println(err)
		return
	}

	username, err := redis.String(conn.Do("GET", "key1"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get key1: %v \n", username)
	}
}
