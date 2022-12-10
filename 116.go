package main

import "fmt"


type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}


func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		nums := []int{}
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]

			if node != nil {
				nums = append(nums, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}

			queue = queue[1:]

			if i < size - 1 && queue[0] != nil {
				node.Next = queue[0]
			}
		}
	}
	return root
}

