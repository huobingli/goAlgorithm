package main

import (
	"fmt"
	"strconv"
)

// Map 映射转化函数
func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, item := range items {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

// Reduce 求和函数
func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, item := range items {
		sum += f(item)
	}
	return sum
}

func main() {
	var users = []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
	}

	ageSlice := mapToString(users, func(user map[string]string) string {
		return user["age"]
	})
	sum := fieldSum(ageSlice, func(age string) int {
		intAge, _ := strconv.Atoi(age)
		return intAge
	})
	fmt.Printf("用户年龄累加结果: %d\n", sum)
}
