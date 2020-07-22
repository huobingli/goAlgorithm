package main

import (
	"context"
	"fmt"
	"time"
)

const DB_ADDRESS = "db_address"
const CALCULATE_VALUE = "calculate_value"

func readDB(ctx context.Context, cost time.Duration) {
	fmt.Println("db address is", ctx.Value(DB_ADDRESS))
	select {
	case <-time.After(cost): //  模拟数据库读取
		fmt.Println("read data from db")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 任务取消的原因
		// 一些清理工作
	}
}
func calculate(ctx context.Context, cost time.Duration) {
	fmt.Println("calculate value is", ctx.Value(CALCULATE_VALUE))
	select {
	case <-time.After(cost): //  模拟数据计算
		fmt.Println("calculate finish")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 任务取消的原因
		// 一些清理工作
	}
}
func main() {
	ctx := context.Background() // 创建一个空的上下文
	// 添加上下文信息
	ctx = context.WithValue(ctx, DB_ADDRESS, "localhost:10086")
	ctx = context.WithValue(ctx, CALCULATE_VALUE, 1234)
	// 设定子 Context 2s 后执行超时返回
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	// 设定执行时间为 4 s
	go readDB(ctx, time.Second*4)
	go calculate(ctx, time.Second*4)

	// 充分执行
	time.Sleep(time.Second * 5)
}
