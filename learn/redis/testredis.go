package main

import (
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