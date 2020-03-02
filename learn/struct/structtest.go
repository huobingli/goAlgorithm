package main

import (
	"fmt"
)

type Bag struct {
	items []int
}

//Insert() 函数将 *Bag 参数放在第一位，强调 Insert 会操作 *Bag 结构体，
//但实际使用中，并不是每个人都会习惯将操作对象放在首位，一定程度上让代码失去一些范式和描述性。
//同时，Insert() 函数也与 Bag 没有任何归属概念，
//随着类似 Insert() 的函数越来越多，面向过程的代码描述对象方法概念会越来越麻烦和难以理解。
// 将一个物品放入背包的过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 类成员函数?
// Insert(itemid int) 的写法与函数一致，(b*Bag) 表示接收器，即 Insert 作用的对象实例。
func (b *Bag) Insert2(itemid int) {
	b.items = append(b.items, itemid)
}

func main() {
	bag := new(Bag)
	Insert(bag, 1)

	bag.Insert2(2)

	for _, value := range bag.items {
		fmt.Printf("%d \n", value)
	}
}
