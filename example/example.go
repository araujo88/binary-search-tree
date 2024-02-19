package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	if n == nil {
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.Value {
		return n.Left.Search(value)
	} else if value > n.Value {
		return n.Right.Search(value)
	}

	return true
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)

	fmt.Println("Contains 3:", root.Search(3))
	fmt.Println("Contains 10:", root.Search(10))
	fmt.Println("Contains 99:", root.Search(99))
}
