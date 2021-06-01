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

	_, err = conn.Do("HSET", "student", "age", 22)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Int64(conn.Do("HGET", "student", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %d \n", res)
	}
}
