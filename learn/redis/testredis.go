package main

import (
	"fmt"
	// 导入redigo扩展包
	"github.com/garyburd/redigo/redis" //go get github.com/garyburd/redigo/redis
)

func main() {
	setPasswd := redis.DialPassword("111111")//设置密码
	// tcp连接redis
	rs, err := redis.Dial("tcp", "47.114.171.118:6380", setPasswd)
	// 操作完后自动关闭
	defer rs.Close()
	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return
	}

	fmt.Println("redis conneted!!")
}
