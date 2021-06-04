package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	setPasswd := redis.DialPassword("111111") //设置密码
	// tcp连接redis
	conn, err := redis.Dial("tcp", "47.114.171.118:6999", setPasswd)
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", "youmen", "18")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	// name, err := redis.String(conn.Do("GET", "youmen"))
	// if err != nil {
	// 	fmt.Println("redis get error:", err)
	// } else {
	// 	fmt.Printf("Get name: %s \n", name)
	// }
}
