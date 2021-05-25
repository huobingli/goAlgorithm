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

	key := "Bonus:start_urls"
	value := "www.10jqka.com.cn"
	// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
	// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	n, err := conn.Do("SETNX", key, value)
	// 若操作失败则返回
	if err != nil {
		fmt.Println(err)
		return
	}
	// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
	if n == int64(1) {
		// 设置过期时间为24小时
		n, _ := conn.Do("EXPIRE", key, 24*3600)
		if n == int64(1) {
			fmt.Println("success")
		}
	} else if n == int64(0) {
		fmt.Println("the key has already existed")
	}
}
