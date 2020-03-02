package main

import (
	"fmt"
)

type Bag struct {
	items []int
}

// 将一个物品放入背包的过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}
func main() {
	bag := new(Bag)
	Insert(bag, 1001)

	for item := bag.items {
		fmt.Printf("%d" item)
	}
}
