package main

type Node struct {
	key   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (tree *BST) Insert(key int) {
	if tree.root == nil {
		tree.root = &Node{key, nil, nil}
	} else {
		currNode := tree.root
		for true {
			if key > currNode.key {
				if currNode.right == nil {
					currNode.right = &Node{key, nil, nil}
					break
				}
				currNode = currNode.right
			} else if key < currNode.key {
				if currNode.left == nil {
					currNode.left = &Node{key, nil, nil}
					break
				}
				currNode = currNode.left
			} else {
				return
			}
		}
	}
}

func (tree *BST) Find(key int) *Node {
	if tree.root == nil {
		return nil
	}
	currNode := tree.root
	for true {
		if key > currNode.key {
			if currNode.right == nil {
				return nil
			}
			currNode = currNode.right
		} else if key < currNode.key {
			if currNode.left == nil {
				return nil
			}
			currNode = currNode.left
		} else {
			return currNode
		}
	}
	return nil
}

func (tree *BST) Remove(key int) {
	if tree.root == nil {
		return
	}
	if tree.root.key == key {
		tree.root = nil
		return
	}
	currNode := tree.root
	for true {
		if key > currNode.key {
			if currNode.right == nil {
				return
			}
			if currNode.right.key == key {
				currNode.right = nil
				return
			}
			currNode = currNode.right
		} else if key < currNode.key {
			if currNode.left == nil {
				return
			}
			if currNode.left.key == key {
				currNode.left = nil
				return
			}
			currNode = currNode.left
		}
	}
}

func (tree *BST) Depth(key int) int {
	if tree.root == nil {
		return 0
	}
	res := 1
	currNode := tree.root
	for true {
		if key > currNode.key {
			if currNode.right == nil {
				return res
			}
			currNode = currNode.right
			res += 1
		} else if key < currNode.key {
			if currNode.left == nil {
				return res
			}
			currNode = currNode.left
			res += 1
		} else {
			return res
		}
	}
	return res
}
