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
	_, err = conn.Do("LPUSH", "list1", "l1", "l2", "l3")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.String(conn.Do("LPOP", "list1"))
	if err != nil {
		fmt.Println("redis POP error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %s \n", res)
	}
}
