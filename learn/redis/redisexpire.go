package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	setPasswd := redis.DialPassword("111111") //设置密码
	// tcp连接redis
	conn, err := redis.Dial("tcp", "47.114.171.118:6380", setPasswd)
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("SET", "name", "youmen")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	_, err = conn.Do("expire", "name", 5)
	if err != nil {
		fmt.Println("set expire error:", err)
		return
	}

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("GET name: %s \n", name)
	}

	time.Sleep(6 * 1000)
	name1, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("GET name: %s \n", name1)
	}
}
