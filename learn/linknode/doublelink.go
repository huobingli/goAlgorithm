package main

import "fmt"

type Node struct {
	data  int
	next  *Node
	front *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		//fmt.Println(*p)
		fmt.Printf("%+v \n", *p)
		p = p.next //移动指针
	}
}
func main() {
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点

	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		(*tail).next = &node
		node.front = tail
		tail = &node
	}
	Shownode(head) //遍历结果
}
