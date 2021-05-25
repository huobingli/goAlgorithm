package main

import (
	"fmt"
	// 导入redigo扩展包
	"github.com/garyburd/redigo/redis" //go get github.com/garyburd/redigo/redis
)

func main() {
	// tcp连接redis
	rs, err := redis.Dial("tcp", "47.114.171.118:6380")
	// 操作完后自动关闭
	defer rs.Close()
	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return
	}

	// 选择db
	//	rs.Do("SELECT", db)

	key := "aaa"
	value := "bbb"
	// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
	// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	n, err := rs.Do("SETNX", key, value)
	// 若操作失败则返回
	if err != nil {
		fmt.Println(err)
		return
	}
	// 返回的n的类型是int64的，所以得将1或0转换成为int64类型的再比较
	if n == int64(1) {
		// 设置过期时间为24小时
		n, _ := rs.Do("EXPIRE", key, 24*3600)
		if n == int64(1) {
			fmt.Println("success")
		}
	} else if n == int64(0) {
		fmt.Println("the key has already existed")
	}
}
