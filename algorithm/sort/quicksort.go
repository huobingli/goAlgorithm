// 快速排序
package main

import . "fmt"

func echosort(sort []int) {
	for i := 0; i < len(sort); i++ {
		Printf("%d ", sort[i])
	}

	Printf("\n")
}

// 每次取出一个归位
// 然后左右开始递归遍历

func quicksort(sort []int, begin int, end int) {
	i := begin
	j := end
	if i < j {
		tmp := sort[i] // 需要归位的元素
		for i != j {
			for i < j && sort[j] > tmp { // 尾部开始找大于归位的元素
				j--
			}
			sort[i] = sort[j]
			for i < j && sort[i] < tmp { // 头部开始找小于归位的元素
				i++
			}
			sort[j] = sort[i]
		}
		sort[i] = tmp
		echosort(sort) // 每次归位输出，查看归位进度
		quicksort(sort, begin, i-1)
		quicksort(sort, i+1, end)
	}
}

func main() {
	//fmt.Println("Hello, world!")
	sort := []int{6, 8, 7, 9, 0, 1, 3, 2, 4, 5}
	quicksort(sort, 0, 9)

	echosort(sort)
}
