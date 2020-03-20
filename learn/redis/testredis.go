package main

import (
<<<<<<< HEAD
"github.com/garyburd/redigo/redis"
"fmt"
)

func main()  {
    conn,err := redis.Dial("tcp","47.114.171.118:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
	}
    defer conn.Close()
}
=======
	"fmt"
	// 导入redigo扩展包
	"github.com/garyburd/redigo/redis" //go get github.com/garyburd/redigo/redis
)

func main() {
	// tcp连接redis
	rs, err := redis.Dial("tcp", "47.114.171.118:6379")
	// 操作完后自动关闭
	defer rs.Close()
	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return
	}
}
>>>>>>> 60bef4fde156596249aeaacbfba5b863fc1fcd75
