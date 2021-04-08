package main
import (
	"strconv"
)

func main() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1)  // 将字符串转化为整型，v2 = 100
	println(v2)

	v3 := 100
	v4 := strconv.Itoa(v3)   // 将整型转化为字符串, v4 = "100"
	println(v4)

	v5 := "true"
	v6, _ := strconv.ParseBool(v5)  // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)  // 将布尔值转化为字符串

	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64)   // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)   // 将整型转化为字符串，第二个参数表示进制

	v9, _ := strconv.ParseUint(v7, 10, 64)   // 将字符串转化为无符号整型，参数含义同 ParseInt
	v7 = strconv.FormatUint(v9, 10)  // 将字符串转化为无符号整型，参数含义同 FormatInt

	v10 := "99.99"
	v11, _ := strconv.ParseFloat(v10, 64)   // 将字符串转化为浮点型，第二个参数表示精度
	v10 = strconv.FormatFloat(v11, 'E', -1, 64)

	q := strconv.Quote("Hello, 世界")    // 为字符串加引号
	q = strconv.QuoteToASCII("Hello, 世界")  // 将字符串转化为 ASCII 编码
	println(q)
}