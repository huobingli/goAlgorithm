package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

func main() {
	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}

	//从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。
	//sort.Slice() 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func(i,j int)bool
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
