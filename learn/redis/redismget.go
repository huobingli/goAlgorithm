package main

import (
	"fmt"
	"reflect"

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

	_, err = conn.Do("MSET", "name", "youmen", "age", "22")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Strings(conn.Do("MGET", "name", "age"))
	if err != nil {
		fmt.Println("redis get error", err)
	} else {
		res_type := reflect.TypeOf(recover())
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("MGET name: %s \n", res)
		fmt.Println(len(res))
	}
}
