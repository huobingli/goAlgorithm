package main

import (
	"fmt"
)


type Node struct {
	data int
	left *Node
	right *Node
 }

func PreOrder(root *Node){
	//a := 1
	if (root != nil){
		fmt.Printf("%d ", root.data)
		PreOrder(root.left)
		PreOrder(root.right)
	}	
}

func InOrder(root *Node){
	if(root!=nil){
		InOrder(root.left)
		fmt.Printf("%d ", root.data)
		InOrder(root.right)
	}
}

func PostOrder(root *Node){
	if(root!=nil){
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func CreateTree() *Node{
	var root *Node = new(Node)		// 根节点

	var left *Node = new(Node)
	var right *Node = new(Node)
	root.left = left
	root.right = right
	root.data = 1

	var left1 *Node = new(Node)
	var right1 *Node = new(Node)
	root.left.left = left1
	root.left.right = right1
	root.left.data = 2
	left1.data = 3
	left1.right = nil
	left1.left = nil
	right1.data = 4
	right1.right = nil
	right1.left = nil

	var left2 *Node = new(Node)
	var right2 *Node = new(Node)
	root.right.left = left2
	root.right.right = right2
	root.right.data = 5
	left2.data = 6
	left2.right = nil
	left2.left = nil
	right2.data = 7
	right2.right = nil
	right2.left = nil

	return root
}

/*		构造的树
				1
			/		\
		2				5
	/		\		/		\
	3		4		6		7
*/

func main(){
	var root *Node = new(Node)
	root = CreateTree()

	fmt.Printf("\nPreOrder------\n")
	PreOrder(root)
	fmt.Printf("\nInOrder------\n")
	InOrder(root)
	fmt.Printf("\nPostOrder------\n")
	PostOrder(root)
}