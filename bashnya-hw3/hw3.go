package main

import (
	"fmt"
)

func main() {
	tree := BST{}
	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(6)
	eightNode := tree.Find(8)
	fmt.Println(*eightNode)
	fmt.Println(eightNode.left.key) // must be 6
}
